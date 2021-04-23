package services

import (
	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/ducnt114/go-gin-demo/utils"
)

type ServiceProvider interface {
	GetAuthService() AuthService
}

type serviceProviderImpl struct {
	authService AuthService
}

func NewServiceProvider(repoProvider repositories.RepositoryProvider, jwtHelper utils.JWTHelper) ServiceProvider {
	authService := newAuthService(repoProvider.GetUserRepo(),
		repoProvider.GetTokenRepo(),
		jwtHelper)

	return &serviceProviderImpl{
		authService: authService,
	}
}

func (s serviceProviderImpl) GetAuthService() AuthService {
	return s.authService
}
