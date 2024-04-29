package services

import (
	"goRepositoryPattern/database/models"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/validators"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error) {
	// validate user input
	if err := arg.Validate(); err != nil {
		return models.RegisterResponse{}, err
	}

	// call register repository
	user, err := s.repo.Register(ctx, arg)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error) {
	// Validate request parameters
	if err := arg.Validate(); err != nil {
		return models.LoginResponse{}, err
	}

	userResponse, err := s.repo.Login(ctx, arg)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return userResponse, nil
}

func (s AuthService) ResendRegistrationOtp(ctx *gin.Context, arg validators.ResendRegistrationOtpInput) (models.ResendRegistrationOtpResponse, error) {

	user, err := s.repo.ResendRegistrationOTP(ctx, arg)
	if err != nil {
		return models.ResendRegistrationOtpResponse{}, err
	}

	return user, nil
	// a, err := s.repo.GetAccount("accounts.email = ?", i.Email)
	// if err != nil {
	// 	R.Status = true
	// 	ctx.JSON(c.Response(http.StatusOK, R))
	// 	return
	// }

	// // Check if email is verified
	// _, err = c.repo.GetAccountToken("account_tokens.type = ? AND accounts.id = ?", uint(constants.AccountTokenTypeEmailConfirmationKey), a.ID)
	// if err == gorm.ErrRecordNotFound {
	// 	R.Message = messages.EmailIsVerified
	// 	ctx.JSON(c.Response(http.StatusBadRequest, R))
	// 	return
	// }

	// t := utils.GenerateRandomNumber(4)

	// c.repo.UpdateAccountToken(models.AccountToken{
	// 	Token: t,
	// }, "account_id = ? AND type = ?", a.ID, uint(constants.AccountTokenTypeEmailConfirmationKey))

	// // Resend registration token
	// go func() {
	// 	tasks.RegisterOtpTask(tasks.RegisterOtpInput{
	// 		Email:     a.Email,
	// 		FirstName: a.FirstName,
	// 		Token:     t,
	// 	})
	// }()

	// ctx.JSON(c.Response(http.StatusOK, R))
}
