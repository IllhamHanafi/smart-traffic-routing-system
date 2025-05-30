package internal

import (
	"regexp"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/errorwrapper"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/ginwrapper"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ProcessRegisterUser(c *gin.Context, input model.RegisterUserInput) {
	errValidation := s.IsRegisterUserRequestValid(input)
	if errValidation.IsError() {
		ginwrapper.RespondWithError(c, errValidation)
		return
	}

	// encrypt password
	encryptedPassword, err := s.encryptPassword(input.Password)
	if err != nil {
		ginwrapper.RespondWithError(c, errorwrapper.ErrInternalServerError)
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
		ginwrapper.RespondWithError(c, model.ErrCreateUser)
		return
	}

	ginwrapper.RespondWithSuccess(c, model.RegisterUserResponse{
		UserID: id,
	})
}

func (s *Service) encryptPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.config.BcryptPasswordRound)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

func (s *Service) IsRegisterUserRequestValid(input model.RegisterUserInput) errorwrapper.ErrorWrapper {
	// validate name
	if input.Name == "" {
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "name is required"})
	}

	// validate role
	if input.Role == "" {
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "role is required"})
	}

	// validate email
	if input.Email == "" {
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "email is required"})
	}
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input.Email)
	if err != nil || !match {
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "email is invalid"})
	}

	// validate password
	// min 8 char
	// min 1 uppercase
	// min 1 number
	// min 1 special char
	if input.Password == "" {
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "password is required"})
	}
	if len(input.Password) < 8 {
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "password is too short"})
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
		return model.ErrInvalidInput.WithDetail(map[string]any{"err": "password is invalid"})
	}

	return errorwrapper.ErrorWrapper{}
}
