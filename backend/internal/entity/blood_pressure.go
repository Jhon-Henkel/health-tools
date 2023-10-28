package entity

import (
	"errors"
	"time"

	"github.com/Jhon-Henkel/health-tools/tree/main/backend/pkg/entity"
)

type BloodPressure struct {
	ID entity.ID `json:"id"`
	Systolic int `json:"systolic"`
	Diastolic int `json:"diastolic"`
	Pulse int `json:"pulse"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrSystolicRequired = errors.New("systolic is required")
	ErrDiastolicRequired = errors.New("diastolic is required")
	ErrPulseRequired = errors.New("pulse is required")
)

func NewBloodPressure(systolic int, diastolic int, pulse int) (*BloodPressure, error) {
	bloodPressure := &BloodPressure{
		ID: entity.NewID(),
		Systolic: systolic,
		Diastolic: diastolic,
		Pulse: pulse,
		CreatedAt: time.Now(),
	}
	err := bloodPressure.Validate()
	if err != nil {
		return nil, err
	}
	return bloodPressure, nil
}

func (b *BloodPressure) Validate() error {
	if b.Systolic == 0 {
		return ErrSystolicRequired
	}
	if b.Diastolic == 0 {
		return ErrDiastolicRequired
	}
	if b.Pulse == 0 {
		return ErrPulseRequired
	}
	if _, err := entity.ParseID(b.ID.String()); err != nil {
		return ErrIdInvalid
	}
	return nil
}