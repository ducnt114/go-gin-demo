package repositories

import "github.com/jinzhu/gorm"

type RepositoryProvider interface {
	GetUserRepo() UserRepository
	GetTokenRepo() TokenRepository
}

type repoProviderImpl struct {
	userRepo  UserRepository
	tokenRepo TokenRepository
}

func NewRepositoryProvider(db *gorm.DB) (RepositoryProvider, error) {
	userRepo := newUserRepository(db)
	tokenRepo := newTokenRepository(db)

	return &repoProviderImpl{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}, nil
}

func (r *repoProviderImpl) GetUserRepo() UserRepository {
	return r.userRepo
}

func (r *repoProviderImpl) GetTokenRepo() TokenRepository {
	return r.tokenRepo
}
