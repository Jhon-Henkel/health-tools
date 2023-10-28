package dto

type BloodPressureInput struct {
	Systolic  int `json:"systolic"`
	Diastolic int `json:"diastolic"`
	Pulse     int `json:"pulse"`
}
