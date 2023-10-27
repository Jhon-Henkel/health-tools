<?php

namespace App\Models\BloodPressure;

use App\Enum\DateEnum;
use Illuminate\Database\Eloquent\Model;

class BloodPressure extends Model
{
    protected $table = 'blood_pressure';
    protected $fillable = ['id', 'systolic', 'diastolic', 'pulse'];
    protected $casts = ['created_at' => DateEnum::MODEL_DEFAULT_DATE_FORMAT];
    protected $hidden = [];
    public $timestamps = true;
}