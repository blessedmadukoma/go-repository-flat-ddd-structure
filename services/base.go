package services

type Service struct {
	AuthService    *AuthService
	UserService    *UserService
	AccountService *AccountsService
	// Add more services as needed
}

func NewService(authService AuthService, userService UserService, accountService AccountsService) *Service {
	return &Service{
		AuthService:    &authService,
		UserService:    &userService,
		AccountService: &accountService,
	}
}
