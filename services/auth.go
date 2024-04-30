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
	if err := arg.Validate(); err != nil {
		return models.RegisterResponse{}, err
	}

	user, err := s.repo.Register(ctx, arg)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error) {

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
}

func (s AuthService) VerifyAccount(ctx *gin.Context, arg validators.VerifyAccountInput) (models.VerifyAccountResponse, error) {

	user, err := s.repo.VerifyAccount(ctx, arg)
	if err != nil {
		return models.VerifyAccountResponse{}, err
	}

	return user, nil
}

func (s AuthService) PasswordReset(ctx *gin.Context, arg validators.PasswordResetInput) (models.ForgotPasswordResponse, error) {

	response, err := s.repo.PasswordReset(ctx, arg)
	if err != nil {
		return models.ForgotPasswordResponse{}, err
	}

	return response, nil
}

func (s AuthService) PasswordResetConfirm(ctx *gin.Context, arg validators.PasswordResetConfirmInput) (models.ForgotPasswordConfirmResponse, error) {
	response, err := s.repo.PasswordResetConfirm(ctx, arg)
	if err != nil {
		return models.ForgotPasswordConfirmResponse{}, err
	}

	return response, nil
}
