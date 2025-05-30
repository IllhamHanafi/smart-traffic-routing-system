package internal

import (
	"time"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/errorwrapper"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/ginwrapper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ProcessLoginUser(c *gin.Context, input model.LoginUserInput) {
	res, err := s.repository.GetUserByEmail(c, input.Email)
	if err != nil {
		ginwrapper.RespondWithError(c, model.ErrInvalidCredentials)
		return
	}

	err = s.checkPasswordIsCorrect(res.Password, input.Password)
	if err != nil {
		ginwrapper.RespondWithError(c, model.ErrInvalidCredentials)
		return
	}

	token, err := s.generateTokenFromUser(res)
	if err != nil {
		ginwrapper.RespondWithError(c, errorwrapper.ErrInternalServerError)
		return
	}

	ginwrapper.RespondWithSuccess(c, model.LoginUserResponse{
		AccessToken: token,
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
