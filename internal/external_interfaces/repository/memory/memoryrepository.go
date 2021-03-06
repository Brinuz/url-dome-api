package repository

// MemoryRepository implements Repository interface
type MemoryRepository struct {
	hashes map[string]string
}

// New returns a valid instace of MemoryRepository
func New() *MemoryRepository {
	return &MemoryRepository{
		hashes: make(map[string]string),
	}
}

// Save saves into memory the current url and hash
func (r *MemoryRepository) Save(url, hash string) error {
	r.hashes[hash] = url
	return nil
}

// Find looks in the memory the current hash and returns matching url
func (r MemoryRepository) Find(hash string) string {
	return r.hashes[hash]
}

// Count returns the amount of entries in memory
func (r MemoryRepository) Count() int {
	return len(r.hashes)
}
