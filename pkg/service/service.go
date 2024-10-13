package service

import (
	"github.com/Ivlay/upstore"
	htmlParser "github.com/Ivlay/upstore/pkg/parser"
	"github.com/Ivlay/upstore/pkg/repository"
	"go.uber.org/zap"
)

type User interface {
	FindOrCreate(user upstore.User) (int, error)
	GetAll() ([]upstore.User, error)
}

type Product interface {
	Prepare()
	Update() (int, error)
	Get() (upstore.Product, error)
}

type Service struct {
	User
	Product
}

func New(logger *zap.Logger, repos *repository.Repository, parser *htmlParser.HtmlParser) *Service {
	logger = logger.Named("service")

	logger.Info("Service started")

	UserService := NewUserService(repos.User, logger)
	ProductService := NewProductService(repos.Product, parser, logger)

	return &Service{User: UserService, Product: ProductService}
}
