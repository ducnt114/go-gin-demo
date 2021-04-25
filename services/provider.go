package services

import (
	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/ducnt114/go-gin-demo/utils"
)

type ServiceProvider interface {
	GetAuthService() AuthService
	GetUserService() UserService
}

type serviceProviderImpl struct {
	authService AuthService
	userService UserService
}

func NewServiceProvider(repoProvider repositories.RepositoryProvider, jwtHelper utils.JWTHelper) ServiceProvider {
	authService := newAuthService(repoProvider.GetUserRepo(),
		repoProvider.GetTokenRepo(),
		jwtHelper)
	userService := newUserService(repoProvider.GetUserRepo())

	return &serviceProviderImpl{
		authService: authService,
		userService: userService,
	}
}

func (s serviceProviderImpl) GetAuthService() AuthService {
	return s.authService
}

func (s serviceProviderImpl) GetUserService() UserService {
	return s.userService
}
