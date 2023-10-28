package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jhon-Henkel/health-tools/tree/main/backend/internal/dto"
	"github.com/Jhon-Henkel/health-tools/tree/main/backend/internal/entity"
	"github.com/Jhon-Henkel/health-tools/tree/main/backend/internal/infra/database"
	"github.com/go-chi/chi/v5"
)

type BloodPressureHandler struct {
	BloodPressureDB database.BloodPressureInterface
}

type Error struct {
	Message string `json:"message"`
}

func NewBloodPressureHandler(db database.BloodPressureInterface) *BloodPressureHandler {
	return &BloodPressureHandler{BloodPressureDB: db}
}

func (b *BloodPressureHandler) CreateBloodPressure(w http.ResponseWriter, r *http.Request) {
	var bloodPressure dto.BloodPressureInput
	err := json.NewDecoder(r.Body).Decode(&bloodPressure)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	dataToInsert, err := entity.NewBloodPressure(bloodPressure.Systolic, bloodPressure.Diastolic, bloodPressure.Pulse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = b.BloodPressureDB.Create(dataToInsert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (b *BloodPressureHandler) GetBloodPressureList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")
	list, err := b.BloodPressureDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func (b *BloodPressureHandler) GetBloodPressure(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}
	item, err := b.BloodPressureDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (b *BloodPressureHandler) DeleteBloodPressure(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err := b.BloodPressureDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = b.BloodPressureDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
