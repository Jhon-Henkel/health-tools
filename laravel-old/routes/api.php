<?php

use App\Enum\RouteEnum;
use App\Http\Controllers\BloodPressure\BloodPressureController;
use Illuminate\Support\Facades\Route;

/** @var Route $router */
$router->prefix('/')->group(function ($router) {
    $router->prefix('blood-pressure')->group(function () use ($router) {
        $router->get('', function () {dd("ssss");})->name(RouteEnum::API_BLOOD_PRESSURE_INDEX);
    });
//    $router->prefix('blood-glucose')->group(function () use ($router) {
//
//    });
});