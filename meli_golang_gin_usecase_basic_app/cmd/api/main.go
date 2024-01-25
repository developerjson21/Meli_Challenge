package main

import (
	"database/sql"
	"fmt"
	"log"
	"meli_golang_gin_basic_app/cmd/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	Start()
}


func Start() {
	
	//* Cargando mis variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	db, errdb := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/db_meli_scanner", os.Getenv("USERDB"), os.Getenv("PASSDB")))
	if errdb != nil {
		log.Fatalln(errdb)
		panic(errdb.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	//Create default gin router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())

	//Set endpoint routes
	ConfigureRoutes(router, db)

	//verify if a port is specified
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := router.Run(":" + port)
	if err != nil {
		panic("could not start app")
	}
	
}
