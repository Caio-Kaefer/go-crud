package main

import (
	"Caio-Kaefer/go-crud/controllers"
	"Caio-Kaefer/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	r := gin.Default()
	r.POST("/createuser", controllers.UsersCreate)
	r.GET("/getusers", controllers.UsersRead)
	r.PUT("/updateuser/:id", controllers.UpdateUser)
	r.DELETE("/deleteuser/:id", controllers.DeleteUser)
	r.Run()
}
