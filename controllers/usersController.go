package controllers

import (
	"Caio-Kaefer/go-crud/initializers"
	"Caio-Kaefer/go-crud/models"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CREATE
func UsersCreate(c *gin.Context) {
	var userInput CreateUserInput
	// erro nos campos
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: userInput.Name, Email: userInput.Email, Password: userInput.Password}
	//inserir no banco
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	//retorno
	c.JSON(200, gin.H{"user": user})
}

// READ
func UsersRead(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	c.JSON(200, gin.H{"users": users})
}

// UPDATE
func UpdateUser(c *gin.Context) {
	//pegar o id pela url
	userID := c.Param("id")
	var existingUser models.User
	result := initializers.DB.First(&existingUser, userID)
	//verificando se o usuário existe
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var userInput CreateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	// Atualizar os campos
	existingUser.Name = userInput.Name
	existingUser.Email = userInput.Email
	existingUser.Password = userInput.Password

	// Salvar as alterações e verificar erros
	if err := initializers.DB.Save(&existingUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	// Retorno
	c.JSON(200, gin.H{"user": existingUser})
}

// DELETE
func DeleteUser(c *gin.Context) {
	//pegar o id pela url
	userID := c.Param("id")

	//verificar se o usuario existe
	var existingUser models.User
	result := initializers.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Excluir o usuário
	if err := initializers.DB.Delete(&existingUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Erro ao excluir usuário"})
		return
	}

	// Retorno
	c.JSON(200, gin.H{"message": "Usuário excluído com sucesso"})
}
