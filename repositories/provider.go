package repositories

import "github.com/jinzhu/gorm"

type RepositoryProvider interface {
	GetUserRepo() UserRepository
}

type repoProviderImpl struct {
	userRepo UserRepository
}

func NewRepositoryProvider(db *gorm.DB) (RepositoryProvider, error) {
	userRepo := newUserRepository(db)

	return &repoProviderImpl{
		userRepo: userRepo,
	}, nil
}

func (r *repoProviderImpl) GetUserRepo() UserRepository {
	return r.userRepo
}
