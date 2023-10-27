<?php

namespace App\Http\Response;

use Illuminate\Http\JsonResponse;

class ResponseError
{
    public static function responseError(string $error, int $statusCode): JsonResponse
    {
        $message = json_decode($error, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            $message = $error;
        }
        return response()->json(['message' => $message], $statusCode);
    }
}