<?php

namespace App\Services\BloodPressure;

use App\Repositories\BloodPressure\BloodPressureRepository;
use App\Services\BaseService;

class BloodPressureService extends BaseService
{
    public function __construct(readonly private BloodPressureRepository $repository)
    {
        // TODO: Implement getRepository() method.
    }

    protected function getRepository(): BloodPressureRepository
    {
        return $this->repository;
    }
}