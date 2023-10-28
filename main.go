package main

import (
	"net/http"

	"github.com/Jhon-Henkel/health-tools/tree/main/configs"
	"github.com/Jhon-Henkel/health-tools/tree/main/internal/entity"
	"github.com/Jhon-Henkel/health-tools/tree/main/internal/infra/database"
	"github.com/Jhon-Henkel/health-tools/tree/main/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dsn := config.DbUser + ":" + config.DbPassword + "@tcp(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.BloodGlucose{}, &entity.BloodPressure{})

	bloodGlucoseDB := database.NewBloodGlucose(db)
	bloodGlucoseHandler := handlers.NewBloodGlucoseHandler(bloodGlucoseDB)

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Route("/blood-glucose", func(router chi.Router) {
		router.Post("/", bloodGlucoseHandler.CreateBloodGlucose)
		router.Get("/", bloodGlucoseHandler.GetBloodGlucoseList)
		router.Get("/{id}", bloodGlucoseHandler.GetBloodGlucose)
		router.Delete("/{id}", bloodGlucoseHandler.DeleteBloodGlucose)
	})

	bloodPressureDB := database.NewBloodPressure(db)
	bloodPressureHandler := handlers.NewBloodPressureHandler(bloodPressureDB)

	router.Route("/blood-pressure", func(router chi.Router) {
		router.Post("/", bloodPressureHandler.CreateBloodPressure)
		router.Get("/", bloodPressureHandler.GetBloodPressureList)
		router.Get("/{id}", bloodPressureHandler.GetBloodPressure)
		router.Delete("/{id}", bloodPressureHandler.DeleteBloodPressure)
	})

	http.ListenAndServe(":8000", router)
}
