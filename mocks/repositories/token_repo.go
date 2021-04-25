package repositories

import (
	"github.com/ducnt114/go-gin-demo/models"
	"github.com/stretchr/testify/mock"
)

type MockTokenRepo struct {
	mock.Mock
}

func (r *MockTokenRepo) Save(m *models.Token) error {
	ret := r.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Token) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (r *MockTokenRepo) Delete(m *models.Token) error {
	ret := r.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Token) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (r *MockTokenRepo) FindByRefreshToken(refreshToken string) (*models.Token, error) {
	ret := r.Called(refreshToken)

	var r0 *models.Token
	if rf, ok := ret.Get(0).(func(string) *models.Token); ok {
		r0 = rf(refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
