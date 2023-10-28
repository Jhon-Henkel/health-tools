package database

import (
	"testing"

	"github.com/Jhon-Henkel/health-tools/tree/main/backend/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestBloodGlucose_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodGlucose{})
	bloodGlocose, err := entity.NewBloodGlucose(100)
	assert.Nil(t, err)
	bloodGlucoseDB := NewBloodGlucose(db)

	err = bloodGlucoseDB.Create(bloodGlocose)
	assert.Nil(t, err)
	assert.NotEmpty(t, bloodGlocose.ID)
}

func TestBloodGlucose_FindAll_withPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodGlucose{})
	for i := 1; i < 24; i++ {
		bloodGlocose, err := entity.NewBloodGlucose(i)
		assert.Nil(t, err)
		db.Create(bloodGlocose)
		assert.Nil(t, err)
	}
	bloodGlocoseDB := NewBloodGlucose(db)
	bloodGlocoseList, err := bloodGlocoseDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(bloodGlocoseList))
	assert.Equal(t, 1, bloodGlocoseList[0].BloodGlucose)
	assert.Equal(t, 10, bloodGlocoseList[9].BloodGlucose)

	bloodGlocoseList, err = bloodGlocoseDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(bloodGlocoseList))
	assert.Equal(t, 11, bloodGlocoseList[0].BloodGlucose)
	assert.Equal(t, 20, bloodGlocoseList[9].BloodGlucose)

	bloodGlocoseList, err = bloodGlocoseDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(bloodGlocoseList))
	assert.Equal(t, 21, bloodGlocoseList[0].BloodGlucose)
	assert.Equal(t, 23, bloodGlocoseList[2].BloodGlucose)
}

func TestBloodGlucose_FindAll_withoutPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodGlucose{})
	for i := 1; i < 24; i++ {
		bloodGlucose, err := entity.NewBloodGlucose(i)
		assert.Nil(t, err)
		db.Create(bloodGlucose)
		assert.Nil(t, err)
	}
	bloodGlucoseDB := NewBloodGlucose(db)
	bloodGlucoseList, err := bloodGlucoseDB.FindAll(0, 0, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 23, len(bloodGlucoseList))
	assert.Equal(t, 1, bloodGlucoseList[0].BloodGlucose)
	assert.Equal(t, 23, bloodGlucoseList[22].BloodGlucose)
}

func TestBloodGlucose_FindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodGlucose{})
	bloodGlucose, err := entity.NewBloodGlucose(100)
	assert.Nil(t, err)
	db.Create(bloodGlucose)
	bloodGlucoseDB := NewBloodGlucose(db)
	bloodGlucoseFound, err := bloodGlucoseDB.FindById(bloodGlucose.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, bloodGlucose.ID, bloodGlucoseFound.ID)
	assert.Equal(t, bloodGlucose.BloodGlucose, bloodGlucoseFound.BloodGlucose)
}

func TestBloodGlucose_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.BloodGlucose{})
	bloodGlucose, err := entity.NewBloodGlucose(100)
	assert.Nil(t, err)
	db.Create(bloodGlucose)
	bloodGlucoseDB := NewBloodGlucose(db)
	err = bloodGlucoseDB.Delete(bloodGlucose.ID.String())
	assert.Nil(t, err)
	_, err = bloodGlucoseDB.FindById(bloodGlucose.ID.String())
	assert.Error(t, err)
}
