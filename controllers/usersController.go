package controllers

//Controller de usuarios
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

type UserResponse struct {
	User models.User `json:"user"`
}

type UserListResponse struct {
	Users []models.User `json:"users"`
}

// UsersCreate cria um usuário
// @Summary Cria um usuário
// @Tags Users
// @Accept json
// @Produce json
// @Param body body CreateUserInput true "Objeto JSON contendo dados do usuário"
// @Success 200 {object} UserResponse
// @Router /api/v1/users/createuser [post]
func UsersCreate(c *gin.Context) {
	var userInput CreateUserInput
	// Verifica os campos
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(userInput.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "erro ao criptografar senha"})
	}

	user := models.User{Name: userInput.Name, Email: userInput.Email, Password: hashedPassword}
	// Insere no banco
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	// Retorna uma mensagem de sucesso
	c.JSON(200, user)
}

// UsersRead retorna uma lista de todos os usuários
// @Summary Retorna uma lista de usuários
// @Tags Users
// @Produce json
// @Success 200 {object} UserListResponse
// @Router /api/v1/users/getusers [get]
func UsersRead(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	response := UserListResponse{
		Users: users,
	}

	c.JSON(200, response)
}

// UpdateUser atualiza as informações de um usuário pelo ID
// @Summary Atualiza um usuário existente
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário a ser atualizado"
// @Param body body CreateUserInput true "Objeto JSON contendo novos dados do usuário"
// @Success 200 {object} UserResponse
// @Router /api/v1/users/updateuser/{id} [put]
func UpdateUser(c *gin.Context) {
	// Pegar o ID pela URL
	userID := c.Param("id")
	var existingUser models.User
	result := initializers.DB.First(&existingUser, userID)
	// Verifica se o usuário existe
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var userInput CreateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	// Atualiza os campos
	existingUser.Name = userInput.Name
	existingUser.Email = userInput.Email
	existingUser.Password = userInput.Password

	// Salva as alterações e verifica erros
	if err := initializers.DB.Save(&existingUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	// Retorna uma mensagem de sucesso
	c.JSON(200, existingUser)
}

// DeleteUser deleta um usuário pelo ID
// @Summary Deleta um usuário pelo ID
// @Tags Users
// @Produce json
// @Param id path string true "ID do usuário a ser excluído"
// @Success 200 {string} string "Usuário excluído com sucesso"
// @Router /api/v1/users/deleteuser/{id} [delete]
func DeleteUser(c *gin.Context) {
	// Pegar o ID pela URL
	userID := c.Param("id")

	// Verifica se o usuário existe
	var existingUser models.User
	result := initializers.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Exclui o usuário
	if err := initializers.DB.Delete(&existingUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Erro ao excluir usuário"})
		return
	}

	// Retorna uma mensagem de sucesso
	c.JSON(200, "Usuário excluído com sucesso")
}
