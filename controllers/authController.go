package controllers

//Controle para a autenticação
import (
	"Caio-Kaefer/go-crud/initializers"
	"Caio-Kaefer/go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func init() {
	initializers.LoadEnvVariables()
}

// UserAuthenticate realiza a autenticação
// @Summary Autenticação
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginInput true "Objeto JSON contendo dados do usuário"
// @Success 200 {string} string "Usuário logado com sucesso"
// @Router /api/v1/auth/login [post]
func Auth(c *gin.Context) {
	var logininput LoginInput
	if err := c.ShouldBindJSON(&logininput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Verifica se o usuário existe no banco de dados
	var user models.User
	result := initializers.DB.Where("email = ?", logininput.Email).First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Verifica se a senha está correta
	if CheckPasswordHash(logininput.Password, user.Password) == false {
		c.JSON(401, gin.H{"error": "Credenciais inválidas"})
		return
	}
	//criação do token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	//adição das claims
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o token JWT"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
