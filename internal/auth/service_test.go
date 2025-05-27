package auth_test

import (
	"restApi/internal/auth"
	"restApi/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initEmail = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepository{})

	email, err := authService.Register(initEmail, "1", "bbska")
	if err != nil {
		t.Fatal(err)
	}

	if email != initEmail {
		t.Fatalf("Email %s do not match %s", email, initEmail)
	}
}