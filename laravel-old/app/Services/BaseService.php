<?php

namespace App\Services;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Illuminate\Support\MessageBag;

abstract class BaseService
{
    abstract protected function getRepository();

    public function findAll()
    {
        return $this->getRepository()->findAll();
    }

    public function findById(int $id)
    {
        return $this->getRepository()->findById($id);
    }

    public function insert($item)
    {
        return $this->getRepository()->insert($item);
    }

    public function update(int $id, $item)
    {
        return $this->getRepository()->update($id, $item);
    }

    public function deleteById(int $id)
    {
        return $this->getRepository()->deleteById($id);
    }

    public function isInvalidRequest(Request $request, array $rules): MessageBag|bool
    {
        $validate = Validator::make($request->all(), $rules, $this->rulesInsertMessages());
        return $validate->fails() ? $validate->errors() : false;
    }

    protected function rulesInsertMessages(): array
    {
        return array(
            'required' => 'O :attribute é obrigatório!',
            'unique' => 'O :attribute já existe!',
            'max' => 'O :attribute não pode ser maior que :max caracteres!',
            'min' => 'O :attribute não pode ser menor que :min caracteres!',
            'int' => 'O :attribute deve ser do tipo int!',
            'string' => 'O :attribute deve ser do tipo string!',
            'decimal' => 'O :attribute deve ser do tipo decimal com o mínimo de 0 casa e máximos de 2 casas!',
            'exists' => 'O :attribute não existe!',
        );
    }
}