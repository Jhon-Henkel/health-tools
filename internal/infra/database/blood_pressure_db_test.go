package database

import (
	"testing"

	"github.com/Jhon-Henkel/health-tools/tree/main/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestBloodPressure_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodPressure{})
	bloodPressure, err := entity.NewBloodPressure(100, 101, 102)
	assert.Nil(t, err)
	bloodPressureDB := NewBloodPressure(db)

	err = bloodPressureDB.Create(bloodPressure)
	assert.Nil(t, err)
	assert.NotEmpty(t, bloodPressure.ID)
	assert.Equal(t, bloodPressure.Systolic, 100)
	assert.Equal(t, bloodPressure.Diastolic, 101)
	assert.Equal(t, bloodPressure.Pulse, 102)
}

func TestBloodPressure_FindAll_withPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodPressure{})
	for i := 1; i < 24; i++ {
		bloodPressure, err := entity.NewBloodPressure(i, i+1, i+2)
		assert.Nil(t, err)
		db.Create(bloodPressure)
		assert.Nil(t, err)
	}
	bloodPressureDB := NewBloodPressure(db)
	bloodPressureList, err := bloodPressureDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(bloodPressureList))
	assert.Equal(t, 1, bloodPressureList[0].Systolic)
	assert.Equal(t, 10, bloodPressureList[9].Systolic)

	bloodPressureList, err = bloodPressureDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(bloodPressureList))
	assert.Equal(t, 11, bloodPressureList[0].Systolic)
	assert.Equal(t, 20, bloodPressureList[9].Systolic)

	bloodPressureList, err = bloodPressureDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(bloodPressureList))
	assert.Equal(t, 21, bloodPressureList[0].Systolic)
	assert.Equal(t, 23, bloodPressureList[2].Systolic)
}

func TestBloodPressure_FindAll_withoutPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodPressure{})
	for i := 1; i < 24; i++ {
		bloodPressure, err := entity.NewBloodPressure(i, i+1, i+2)
		assert.Nil(t, err)
		db.Create(bloodPressure)
		assert.Nil(t, err)
	}
	bloodPressureDB := NewBloodPressure(db)
	bloodPressureList, err := bloodPressureDB.FindAll(0, 0, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 23, len(bloodPressureList))
	assert.Equal(t, 1, bloodPressureList[0].Systolic)
	assert.Equal(t, 23, bloodPressureList[22].Systolic)
}

func TestBloodPressure_FindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodPressure{})
	bloodPressure, err := entity.NewBloodPressure(100, 101, 102)
	assert.Nil(t, err)
	db.Create(bloodPressure)
	assert.Nil(t, err)
	bloodPressureDB := NewBloodPressure(db)

	bloodPressure, err = bloodPressureDB.FindById(bloodPressure.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, bloodPressure.Systolic, 100)
	assert.Equal(t, bloodPressure.Diastolic, 101)
	assert.Equal(t, bloodPressure.Pulse, 102)
}

func TestBloodPressure_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodPressure{})
	bloodPressure, err := entity.NewBloodPressure(100, 101, 102)
	assert.Nil(t, err)
	db.Create(bloodPressure)
	assert.Nil(t, err)
	bloodPressureDB := NewBloodPressure(db)

	err = bloodPressureDB.Delete(bloodPressure.ID.String())
	assert.Nil(t, err)
	_, err = bloodPressureDB.FindById(bloodPressure.ID.String())
	assert.Error(t, err)
}
