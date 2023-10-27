<?php

namespace App\Repositories\BloodPressure;

use App\Models\BloodPressure\BloodPressure;
use App\Repositories\BaseRepository;
use App\Resources\BloodPressure\BloodPressureResource;

class BloodPressureRepository extends BaseRepository
{
    public function __construct(
        readonly private BloodPressure $model,
        readonly private BloodPressureResource $resource
    ) {}

    protected function getModel(): BloodPressure
    {
        return $this->model;
    }

    protected function getResource(): BloodPressureResource
    {
        return $this->resource;
    }
}