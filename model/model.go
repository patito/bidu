package model

import "database/sql"

// Model struct
type Model struct {
	db *sql.DB
}

// New creates a new instance of Model struct
func New(db *sql.DB) *Model {
	return &Model{db: db}
}
