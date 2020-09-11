package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skang28/golang_gin_mongo/controller"
)

//init called before main function
func init() {
	// log error if .env file doesn't exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file")
	}
}

func main() {
	r := gin.Default()
	//group the api requests together
	api := r.Group("/api")
	{
		user := new(controller.AccountController)
		api.POST("/register",user.Register)
		api.POST("/login", user.Login)

	}
	
	r.Run(":5000")
}
