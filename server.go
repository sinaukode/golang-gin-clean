package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-clean-code-gin/config"
	"go-clean-code-gin/person/handler"
	"go-clean-code-gin/person/repository"
	"go-clean-code-gin/person/service"
	"log"
)

func main() {
	port := "3030"

	router := gin.Default()
	db := config.DbConnect()
	defer db.Close()

	//Person
	personRepository := repository.CreatePersonRepository(db)
	personService := service.CreatePersonService(personRepository)
	handler.CreatePersonHandler(router, personService)

	err := router.Run(":" + port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running on port" + port)
}
