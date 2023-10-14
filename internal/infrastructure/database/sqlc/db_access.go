package sqlc

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Database interface {
	Querier
}

type BaseRepository struct {
	DB *Queries
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLDbHandler struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewSQLDbHandler(connPool *pgxpool.Pool) Database {
	return &SQLDbHandler{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
