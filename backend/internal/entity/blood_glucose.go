package entity

import (
	"errors"
	"time"

	"github.com/Jhon-Henkel/health-tools/tree/main/backend/pkg/entity"
)

type BloodGlucose struct {
	ID entity.ID `json:"id"`
	BloodGlucose int `json:"blood_glucose"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrIdInvalid = errors.New("id is invalid")
	ErrBloodGlucoseRequired = errors.New("blood_glucose is required")
)

func NewBloodGlucose(bloodGlucose int) (*BloodGlucose, error) {
	glucose := &BloodGlucose{
		ID: entity.NewID(),
		BloodGlucose: bloodGlucose,
		CreatedAt: time.Now(),
	}
	err := glucose.Validate()
	if err != nil {
		return nil, err
	}
	return glucose, nil
}

func (g *BloodGlucose) Validate() error {
	if g.BloodGlucose == 0 {
		return ErrBloodGlucoseRequired
	}
	if _, err := entity.ParseID(g.ID.String()); err != nil {
		return ErrIdInvalid
	}
	return nil
}