package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
)

type MockJWTHelper struct {
	mock.Mock
}

func (h *MockJWTHelper) GenerateToken(claims jwt.Claims) (string, error) {
	ret := h.Called(claims)

	var r0 string
	if rf, ok := ret.Get(0).(func(jwt.Claims) string); ok {
		r0 = rf(claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(jwt.Claims) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (h *MockJWTHelper) ParseClaims(token string, claims jwt.Claims) error {
	ret := h.Called(token, claims)

	var r0 error
	if rf, ok := ret.Get(1).(func(string, jwt.Claims) error); ok {
		r0 = rf(token, claims)
	} else {
		r0 = ret.Error(1)
	}

	return r0
}
