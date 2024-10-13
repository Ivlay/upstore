package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"go.uber.org/zap"
)

func NewSqLiteDB(logger *zap.Logger, path string) (*sql.DB, error) {
	logger = logger.Named("database")

	database, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := database.Ping(); err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Connected to database", zap.String("path", path))

	if err := prepareDB(database); err != nil {
		logger.Fatal(err.Error())
	}

	return database, nil
}

func prepareDB(db *sql.DB) error {
	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

	user := `CREATE TABLE IF NOT EXISTS users
					(
						firstname TEXT NOT NULL,
						username TEXT NOT NULL,
						chat_id INTEGER NOT NULL,
						user_id INTEGER NOT NULL UNIQUE,
						created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
						PRIMARY KEY (user_id, chat_id)
					)`

	userTable, err := db.Prepare(user)
	if err != nil {
		fmt.Print("THis USER ERROR\n")
		return err
	}
	userTable.Exec()

	product := `CREATE TABLE IF NOT EXISTS products
							(
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								title TEXT NOT NULL UNIQUE,
								price INTEGER NOT NULL,
								old_price INTEGER NOT NULL,
								updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
								price_id TEXT NOT NULL
							)
							`
	productTable, err := db.Prepare(product)
	if err != nil {
		return err
	}
	productTable.Exec()

	trigger := `CREATE TRIGGER trigger_set_timestamp
								BEFORE UPDATE of price ON products
							BEGIN
								UPDATE products SET old_price = old.price, updated_at = CURRENT_TIMESTAMP;
							END;`
	db.Exec(trigger)

	productList := `CREATE TABLE IF NOT EXISTS products_lists
								(
									id INTEGER PRIMARY KEY,
									user_id INTEGER NOT NULL,
									product_id INTEGER NOT NULL,
									created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
									FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
									FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
									UNIQUE (user_id, product_id)
								)`

	productListTable, err := db.Prepare(productList)
	if err != nil {
		return err
	}
	productListTable.Exec()

	return nil
}
