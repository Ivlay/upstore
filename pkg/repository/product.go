package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ivlay/upstore"
)

type ProductSql struct {
	db *sql.DB
}

func NewProductSql(db *sql.DB) *ProductSql {
	return &ProductSql{db: db}
}

func (r *ProductSql) Update(product upstore.Product) (int, error) {
	var count int

	query := fmt.Sprintf(`
		update %s
			set price = :price
			where price_id = :price_id and title = :title and price != :price
			returning id
	`, productsTable)

	if err := r.db.QueryRow(query, product.Price, product.PriceId, product.Title, product.Price).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ProductSql) Prepare(product upstore.Product) error {
	query := fmt.Sprintf(`
		insert into %s (price_id, title, price, old_price)
		values (:price_id, :title, :price, :price)`,
		productsTable,
	)

	_, err := r.db.Exec(query, product.PriceId, product.Title, product.Price, product.Price)

	return err
}

func (r *ProductSql) Count() (int, error) {
	var rowCount int
	query := fmt.Sprintf(`
		select count(*) from %s
	`, productsTable)

	if err := r.db.QueryRow(query).Scan(&rowCount); err != nil {
		return 0, err
	}

	return rowCount, nil
}

func (r *ProductSql) Get() (upstore.Product, error) {
	var product upstore.Product

	query := fmt.Sprintf(`
		select price_id, title, price, old_price from %s
	`, productsTable)

	if err := r.db.QueryRow(query).Scan(&product.PriceId, &product.Title, &product.Price, &product.OldPrice); err != nil {
		return product, err
	}

	return product, nil
}
