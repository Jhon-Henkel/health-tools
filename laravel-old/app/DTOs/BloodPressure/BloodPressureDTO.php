<?php

namespace App\DTOs\BloodPressure;

class BloodPressureDTO
{
    public function __construct(
        public null|int $id,
        public int $systolic,
        public int $diastolic,
        public int $pulse,
        public null|string $created_at
    ) {}

    public function getId(): null|int
    {
        return $this->id;
    }

    public function getSystolic(): int
    {
        return $this->systolic;
    }

    public function getDiastolic(): int
    {
        return $this->diastolic;
    }

    public function getPulse(): int
    {
        return $this->pulse;
    }

    public function getCreatedAt(): null|string
    {
        return $this->created_at;
    }
}