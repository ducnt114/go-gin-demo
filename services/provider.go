package services

import "github.com/ducnt114/go-gin-demo/repositories"

type ServiceProvider interface {
	GetAuthService() AuthService
}

type serviceProviderImpl struct {
	authService AuthService
}

func NewServiceProvider(repoProvider repositories.RepositoryProvider) ServiceProvider {
	authService := newAuthService(repoProvider.GetUserRepo())

	return &serviceProviderImpl{
		authService: authService,
	}
}

func (s serviceProviderImpl) GetAuthService() AuthService {
	return s.authService
}
