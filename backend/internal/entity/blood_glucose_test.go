package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBloodGlucose(t *testing.T) {
	bloodGlocose, err := NewBloodGlucose(100)
	
	assert.Nil(t, err)
	assert.Equal(t, bloodGlocose.BloodGlucose, 100)
	assert.NotEmpty(t, bloodGlocose.ID)
	assert.NotNil(t, bloodGlocose.CreatedAt)
}

func TestNewBloodGlucoseError(t *testing.T) {
	bloodGlocose, err := NewBloodGlucose(0)
	
	assert.NotNil(t, err)
	assert.Nil(t, bloodGlocose)
}