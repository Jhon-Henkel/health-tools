package database

import (
	"github.com/Jhon-Henkel/health-tools/tree/main/backend/internal/entity"
	"gorm.io/gorm"
)

type BloodPressure struct {
	DB *gorm.DB
}

func NewBloodPressure(db *gorm.DB) *BloodPressure {
	return &BloodPressure{DB: db}
}

func (bp *BloodPressure) Create(bloodPressure *entity.BloodPressure) error {
	return bp.DB.Create(bloodPressure).Error
}

func (bp *BloodPressure) FindAll(page, limit int, sort string) ([]entity.BloodPressure, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	var bloodPressure []entity.BloodPressure
	var err error
	if page != 0 && limit != 0 {
		err = bp.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&bloodPressure).Error
	} else {
		err = bp.DB.Order("created_at " + sort).Find(&bloodPressure).Error
	}
	return bloodPressure, err
}

func (bp *BloodPressure) FindById(id string) (*entity.BloodPressure, error) {
	var bloodPressure entity.BloodPressure
	return &bloodPressure, bp.DB.First(&bloodPressure, "id = ?", id).Error
}

func (bp *BloodPressure) Delete(id string) error {
	bloodPressure, err := bp.FindById(id)
	if err != nil {
		return err
	}
	return bp.DB.Delete(bloodPressure).Error
}