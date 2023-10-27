<?php

namespace App\Http\Controllers;

use App\Http\Response\ResponseError;
use Illuminate\Database\QueryException;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\MessageBag;
use Symfony\Component\HttpFoundation\Response as HTTP;

abstract class BaseController extends Controller
{
    abstract protected function rulesInsert(): array;
    abstract protected function rulesUpdate(): array;
    abstract protected function getService();
    abstract protected function getResource();

    public function index(): JsonResponse
    {
        try {
            dd("aaa");
            $find = $this->getService()->findAll();
            $itens = $this->getResource()->arrayDtoToVoItens($find);
            return response()->json($itens, HTTP::HTTP_OK);
        } catch (QueryException $exception) {
            return $this->returnErrorDatabaseConnect();
        }
    }

    public function show(int $id): JsonResponse
    {
        try {
            $find = $this->getService()->findById($id);
            return $find
                ? response()->json($this->getResource()->dtoToVo($find), HTTP::HTTP_OK)
                : response()->json('Registro não encontrado!', HTTP::HTTP_NOT_FOUND);
        } catch (QueryException $exception) {
            return $this->returnErrorDatabaseConnect();
        }
    }

    public function insert(Request $request): JsonResponse
    {
        try {
            $invalid = $this->getService()->isInvalidRequest($request, $this->rulesInsert());
            if ($invalid instanceof MessageBag) {
                return ResponseError::responseError($invalid, HTTP::HTTP_BAD_REQUEST);
            }
            $item = $this->getResource()->arrayToDto($request->json()->all());
            $insert = $this->getService()->insert($item);
            return $insert
                ? response()->json($this->getResource()->dtoToVo($insert), HTTP::HTTP_CREATED)
                : response()->json('Erro ao inserir item.', HTTP::HTTP_INTERNAL_SERVER_ERROR);
        } catch (QueryException $exception) {
            return $this->returnErrorDatabaseConnect();
        }
    }

    public function update(int $id, Request $request): JsonResponse
    {
        try {
            $invalid = $this->getService()->isInvalidRequest($request, $this->rulesUpdate());
            if ($invalid instanceof MessageBag) {
                return ResponseError::responseError($invalid, HTTP::HTTP_BAD_REQUEST);
            }
            $requestItem = $request->json()->all();
            $requestItem['id'] = $id;
            $item = $this->getResource()->arrayToDto($requestItem);
            $updated = $this->getService()->update($id, $item);
            return response()->json($this->getResource()->dtoToVo($updated), HTTP::HTTP_OK);
        } catch (QueryException $exception) {
            return $this->returnErrorDatabaseConnect();
        }
    }

    public function delete(int $id): Response|JsonResponse
    {
        try {
            $this->getService()->deleteById($id);
            return response(null, HTTP::HTTP_OK);
        } catch (QueryException $exception) {
            return $this->returnErrorDatabaseConnect();
        }
    }

    protected function returnErrorDatabaseConnect(): JsonResponse
    {
        $message = 'Erro ao se conectar com o banco de dados!';
        return ResponseError::responseError($message, HTTP::HTTP_INTERNAL_SERVER_ERROR);
    }
}