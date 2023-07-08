package db

import (
	//"context"
	"database/sql"

)


// SQLStore provides all functions to execute SQL queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries:  New(db),
	}
}