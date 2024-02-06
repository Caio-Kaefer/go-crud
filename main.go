package main

import (
	"Caio-Kaefer/go-crud/controllers"
	_ "Caio-Kaefer/go-crud/docs"
	"Caio-Kaefer/go-crud/initializers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

//@title GO-CRUD
//@version 1
//@Description Desafio Tecnico para Digital Sys

//@contact.name Caio Kaefer
//@contact.email kaefer.caio@gmail.com

//@host localhost:3000
//BasePath /api/v1

func main() {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	user := v1.Group("/users")
	{
		//CREATE
		user.POST("/createuser", controllers.UsersCreate)

		//READ
		user.GET("/getusers", controllers.UsersRead)

		// UPDATE
		user.PUT("/updateuser/:id", controllers.UpdateUser)

		// DELETE
		user.DELETE("/deleteuser/:id", controllers.DeleteUser)
	}

	auth := v1.Group("/auth")

	{
		//LOGIN
		auth.POST("/login", controllers.Auth)
	}

	// Rota para a documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
