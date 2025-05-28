package internal

import (
	"net/http"
	"time"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ProcessLoginUser(c *gin.Context, input model.LoginUserInput) {
	res, err := s.repository.GetUserByEmail(c, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": model.ErrInvalidCredentials,
			"status":  "error",
		})
		return
	}

	err = s.checkPasswordIsCorrect(res.Password, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": model.ErrInvalidCredentials,
			"status":  "error",
		})
		return
	}

	token, err := s.generateTokenFromUser(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  "success",
		"data": gin.H{
			"token": token,
		},
	})
}

func (s *Service) checkPasswordIsCorrect(hashedPassword string, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

func (s *Service) generateTokenFromUser(user model.User) (string, error) {
	// to do: add cache
	now := time.Now()

	return s.jwt.GenerateAndSignToken(jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     now.Add(s.config.JWT.ExpiredIn),
	})
}
