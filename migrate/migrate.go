package main

import (
	"Caio-Kaefer/go-crud/initializers"
	"Caio-Kaefer/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
