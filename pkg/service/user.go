package service

import (
	"github.com/Ivlay/upstore"
	"github.com/Ivlay/upstore/pkg/repository"
	"go.uber.org/zap"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User, logger *zap.Logger) *UserService {
	return &UserService{repo: repo}
}

func (r *UserService) FindOrCreate(user upstore.User) (int, error) {
	return r.repo.FindOrCreate(user)
}

func (r *UserService) GetAll() ([]upstore.User, error) {
	return r.repo.GetAll()
}
