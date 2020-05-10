package repository

import (
	"database/sql"
	"url-at-minimal-api/internal/domain"
)

// PostgresRepository implements Repository interface
type PostgresRepository struct {
	db *sql.DB
}

// New returns a valid instace of MemoryRepository
func New(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// Save saves into postgres the current url and hash
func (r *PostgresRepository) Save(url, hash string) error {
	_, err := r.db.Exec("INSERT INTO urls (url, hash) VALUES(?, ?)", url, hash)
	if err != nil {
		return domain.ErrCouldNotSaveEntry
	}
	return nil
}

// Find looks in the postgres the current hash and returns matching url
func (r PostgresRepository) Find(hash string) string {
	rows, err := r.db.Query("SELECT url FROM urls WHERE hash = ?", hash)
	if err != nil {
		return ""
	}
	defer rows.Close()
	if !rows.Next() {
		return ""
	}
	var url string
	rows.Scan(&url)
	return url
}
