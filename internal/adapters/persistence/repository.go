package persistence

// Repository interface
type Repository interface {
	Save(url, hash string)
	Find(hash string) string
}
