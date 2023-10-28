package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jhon-Henkel/health-tools/tree/main/internal/dto"
	"github.com/Jhon-Henkel/health-tools/tree/main/internal/entity"
	"github.com/Jhon-Henkel/health-tools/tree/main/internal/infra/database"
	"github.com/go-chi/chi/v5"
)

type BloodGlucoseHandler struct {
	BloodGlucoseDB database.BloodGlucoseInterface
}

func NewBloodGlucoseHandler(db database.BloodGlucoseInterface) *BloodGlucoseHandler {
	return &BloodGlucoseHandler{BloodGlucoseDB: db}
}

func (b *BloodGlucoseHandler) CreateBloodGlucose(w http.ResponseWriter, r *http.Request) {
	var bloodGlucose dto.BloodGlucoseInput
	err := json.NewDecoder(r.Body).Decode(&bloodGlucose)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	dataToInsert, err := entity.NewBloodGlucose(bloodGlucose.BloodGlucose)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = b.BloodGlucoseDB.Create(dataToInsert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (b *BloodGlucoseHandler) GetBloodGlucoseList(w http.ResponseWriter, r *http.Request) {
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
	list, err := b.BloodGlucoseDB.FindAll(pageInt, limitInt, sort)
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

func (b *BloodGlucoseHandler) GetBloodGlucose(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}
	item, err := b.BloodGlucoseDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (b *BloodGlucoseHandler) DeleteBloodGlucose(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err := b.BloodGlucoseDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = b.BloodGlucoseDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
