package database

import "github.com/Jhon-Henkel/health-tools/tree/main/backend/internal/entity"

type BloodPressureInterface interface {
	Create(bloodPressure *entity.BloodPressure) (*entity.BloodPressure, error)
	FindAll(page, limit int, sort string) ([]*entity.BloodPressure, error)
	FindById(id string) (*entity.BloodPressure, error)
	Delete(id string) error
}

type BloodGlucoseInterface interface {
	Create(bloodGlucose *entity.BloodGlucose) (*entity.BloodGlucose, error)
	FindAll(page, limit int, sort string) ([]*entity.BloodGlucose, error)
	FindById(id string) (*entity.BloodGlucose, error)
	Delete(id string) error
}