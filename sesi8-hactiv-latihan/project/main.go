package main

import (
	"fmt"
	"log"
	"project-sesi8/controllers"
	"project-sesi8/models"
	"project-sesi8/repositories"
	"project-sesi8/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	DB_HOST  = "localhost"
	DB_PORT  = "5432"
	DB_USER  = "root"
	DB_PASS  = "root"
	DB_NAME  = "db_go_sql"
	APP_PORT = ":8888"
)

func main() {

	db := connectDB()

	router := gin.Default()

	personRepo := repositories.NewPersonRepo(db)
	personService := services.NewPersonService(personRepo)
	personController := controllers.NewPersonController(personService)

	router.POST("/persons", personController.CreateNewPerson)
	router.GET("/persons", personController.GetAllPerson)
	router.GET("persons/:personId", personController.GetPersonById)
	router.DELETE("persons/:personId", personController.DeletePersonById)
	router.PUT("/persons", personController.UpdatePersonById)

	log.Println("server running at port ", APP_PORT)
	router.Run(APP_PORT)
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	log.Default().Println("connection db success")

	db.AutoMigrate(models.Person{})
	return db
}
