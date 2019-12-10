package repository

import (
	"github.com/go-redis/redis"
)

// RedisRepository implements Repository interface
type RedisRepository struct {
	client *redis.Client
}

// New returns a valid instace of MemoryRepository
func New(c *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: c,
	}
}

// Save saves into memory the current url and hash
func (r *RedisRepository) Save(url, hash string) {
	r.client.Set(hash, url, 0) // Yes, currently ignoring bad state
}

// Find looks in the memory the current hash and returns matching url
func (r RedisRepository) Find(hash string) string {
	return r.client.Get(hash).Val() // Yes, currently ignoring bad state
}

// Count returns the amount of entries in memory
func (r RedisRepository) Count() int64 {
	return r.client.DBSize().Val()
}
