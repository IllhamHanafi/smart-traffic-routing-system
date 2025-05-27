package internal

import (
	"net/http"
	"regexp"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ProcessRegisterUser(c *gin.Context, input model.RegisterUserInput) {
	err := s.IsRegisterUserRequestValid(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "input is invalid",
		})
		return
	}

	// encrypt password
	encryptedPassword, err := s.encryptPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	// insert to database
	id, err := s.repository.CreateUser(c.Request.Context(), repository.CreateUserInput{
		Name:      input.Name,
		Role:      input.Role,
		Email:     input.Email,
		Password:  encryptedPassword,
		CreatedBy: input.CreatedBy,
		UpdatedBy: input.CreatedBy,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  "success",
		"id":      id,
	})
}

func (s *Service) encryptPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.config.BcryptPasswordRound)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

func (s *Service) IsRegisterUserRequestValid(input model.RegisterUserInput) error {
	// validate name
	if input.Name == "" {
		return model.ErrInvalidName
	}

	// validate role
	if input.Role == "" {
		return model.ErrInvalidRole
	}

	// validate email
	if input.Email == "" {
		return model.ErrInvalidEmail
	}
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input.Email)
	if err != nil || !match {
		return model.ErrInvalidEmail
	}

	// validate password
	// min 8 char
	// min 1 uppercase
	// min 1 number
	// min 1 special char
	if input.Password == "" {
		return model.ErrInvalidPassword
	}
	if len(input.Password) < 8 {
		return model.ErrInvalidPassword
	}
	// to do: maybe use regexp to validate password
	hasUppercase := false
	hasNumber := false
	hasSpecialChar := false
	specialCharList := "!@#$%^&*()_+{}|:<>?`-=[]\\;',./"
	for _, rn := range input.Password {
		if rn >= 'A' && rn <= 'Z' {
			hasUppercase = true
		}
		if rn >= '0' && rn <= '9' {
			hasNumber = true
		}
		for _, sc := range specialCharList {
			if rn == sc {
				hasSpecialChar = true
			}
		}
	}

	if !hasUppercase || !hasNumber || !hasSpecialChar {
		return model.ErrInvalidPassword
	}

	return nil
}
