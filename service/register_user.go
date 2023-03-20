package service

import (
	"context"

	"github.com/yukinkoNo10/go_todo_app/entity"
	"github.com/yukinkoNo10/go_todo_app/store"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, name string, password string, role string) (*entity.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}
	if err := r.Repo.UserRegister(ctx, r.DB, u); err != nil {
		return nil, err
	}
	return u, nil
}
