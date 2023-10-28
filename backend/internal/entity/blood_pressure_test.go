package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBloodPressure(t *testing.T) {
	bloodPressure, err := NewBloodPressure(120, 80, 70)
	
	assert.Nil(t, err)
	assert.Equal(t, bloodPressure.Systolic, 120)
	assert.Equal(t, bloodPressure.Diastolic, 80)
	assert.Equal(t, bloodPressure.Pulse, 70)
	assert.NotEmpty(t, bloodPressure.ID)
	assert.NotNil(t, bloodPressure.CreatedAt)
}

func TestBloodPressureSystolicError(t *testing.T) {
	bloodPressure, err := NewBloodPressure(0, 80, 70)
	
	assert.NotNil(t, err)
	assert.Nil(t, bloodPressure)
	assert.Equal(t, ErrSystolicRequired, err)
}

func TestBloodPressureDiastolicError(t *testing.T) {
	bloodPressure, err := NewBloodPressure(120, 0, 70)
	
	assert.NotNil(t, err)
	assert.Nil(t, bloodPressure)
	assert.Equal(t, ErrDiastolicRequired, err)
}

func TestBloodPressurePulseError(t *testing.T) {
	bloodPressure, err := NewBloodPressure(120, 80, 0)
	
	assert.NotNil(t, err)
	assert.Nil(t, bloodPressure)
	assert.Equal(t, ErrPulseRequired, err)
}