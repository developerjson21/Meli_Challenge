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
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)
// @title           API Scanner GitHub Repository
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   Stiven Mesa
// @contact.url    http://www.swagger.io/support
// @contact.email  developerjson21@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
