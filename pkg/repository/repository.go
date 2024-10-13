package repository

import (
	"database/sql"

	"github.com/Ivlay/upstore"
	"go.uber.org/zap"
)

const (
	usersTable         = "users"
	productsTable      = "products"
	productsListsTable = "products_lists"
)

type User interface {
	Create(user upstore.User) (int, error)
	FindOrCreate(user upstore.User) (int, error)
	GetAll() ([]upstore.User, error)
}

type Product interface {
	Prepare(product upstore.Product) error
	Update(product upstore.Product) (int, error)
	Count() (int, error)
	Get() (upstore.Product, error)
}

type Repository struct {
	User
	Product
}

func New(logger *zap.Logger, db *sql.DB) *Repository {
	logger = logger.Named("repository")

	logger.Info("Repository created")

	UserRepository := NewUserSql(db)
	ProductRepository := NewProductSql(db)

	return &Repository{
		User:    UserRepository,
		Product: ProductRepository,
	}
}
