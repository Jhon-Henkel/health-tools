<?php

namespace App\Http\Controllers\BloodPressure;

use App\Exceptions\NotImplementedException;
use App\Http\Controllers\BaseController;
use App\Resources\BloodPressure\BloodPressureResource;
use App\Services\BloodPressure\BloodPressureService;

class BloodPressureController extends BaseController
{
    public function __construct(
        protected BloodPressureService $service,
        protected BloodPressureResource $resource
    ) {}

    protected function rulesInsert(): array
    {
        return [
            'systolic' => 'required|int|max:3',
            'diastolic' => 'required|int|max:3',
            'pulse' => 'required|int|max:3'
        ];
    }

    protected function rulesUpdate(): array
    {
        throw new NotImplementedException();
    }

    protected function getService(): BloodPressureService
    {
        return $this->service;
    }

    protected function getResource(): BloodPressureResource
    {
        return $this->resource;
    }
}