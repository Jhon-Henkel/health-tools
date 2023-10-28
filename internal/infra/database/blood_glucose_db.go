package database

import (
	"github.com/Jhon-Henkel/health-tools/tree/main/internal/entity"
	"gorm.io/gorm"
)

type BloodGlucose struct {
	DB *gorm.DB
}

func NewBloodGlucose(db *gorm.DB) *BloodGlucose {
	return &BloodGlucose{DB: db}
}

func (bg *BloodGlucose) Create(bloodGlucose *entity.BloodGlucose) error {
	return bg.DB.Create(bloodGlucose).Error
}

func (bg *BloodGlucose) FindAll(page, limit int, sort string) ([]entity.BloodGlucose, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	var bloodGlucose []entity.BloodGlucose
	var err error
	if page != 0 && limit != 0 {
		err = bg.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&bloodGlucose).Error
	} else {
		err = bg.DB.Order("created_at " + sort).Find(&bloodGlucose).Error
	}
	return bloodGlucose, err
}

func (bg *BloodGlucose) FindById(id string) (*entity.BloodGlucose, error) {
	var bloodGlucose entity.BloodGlucose
	return &bloodGlucose, bg.DB.First(&bloodGlucose, "id = ?", id).Error
}

func (bg *BloodGlucose) Delete(id string) error {
	bloodGlucose, err := bg.FindById(id)
	if err != nil {
		return err
	}
	return bg.DB.Delete(bloodGlucose).Error
}
