package services

type Service struct {
	AuthService *AuthService
	UserService *UserService
	// Add more services as needed
}

func NewService(authService AuthService, userService UserService) *Service {
	return &Service{
		AuthService: &authService,
		UserService: &userService,
	}
}
