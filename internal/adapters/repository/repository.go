package repository

// Repository interface
type Repository interface {
	Save(url, hash string) error
	Find(hash string) string
}
