package main

import (
	"database/sql"
	"meli_golang_gin_basic_app/cmd/handler"
	"meli_golang_gin_basic_app/cmd/internal/middleware"
	"meli_golang_gin_basic_app/cmd/internal/repository"
	"meli_golang_gin_basic_app/cmd/usecase"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine, DB *sql.DB) {
	

	healtChecker := handler.HealthChecker{}

	router.GET("/ping", healtChecker.PingHandler)

	storageRepository := repository.NewStorageDB(DB)
	dataRepository := repository.NewDataRepository()
	scanUsecase := usecase.NewScanUsecase(dataRepository, storageRepository)
	scanHandler := handler.NewScanHandler(scanUsecase)

	v1 := router.Group("/api/v1")
	{
		//Register scan repo
		v1.POST("/scan", middleware.Authentication(), scanHandler.Scan)

		//Classify database
		v1.GET("/result/:job_id", middleware.Authentication(), scanHandler.GetResult)

	}
}
