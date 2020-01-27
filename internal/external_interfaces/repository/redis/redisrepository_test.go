package repository_test

import (
	"testing"
	"url-at-minimal-api/internal/domain"
	repository "url-at-minimal-api/internal/external_interfaces/repository/redis"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepositorySaveAndFind(t *testing.T) {
	// Given
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	repo := repository.New(client)

	_ = repo.Save("https://www.google.com", "Vsdfb1")
	_ = repo.Save("https://www.microsoft.com", "MasFgr0")

	// When
	count := repo.Count()
	entry1 := repo.Find("Vsdfb1")
	entry2 := repo.Find("MasFgr0")

	// Then
	assert.Equal(t, int64(2), count)
	assert.Equal(t, "https://www.google.com", entry1)
	assert.Equal(t, "https://www.microsoft.com", entry2)
}

func TestRedisRepositoryExistingHash(t *testing.T) {
	// Given
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	repo := repository.New(client)

	// When
	firstSetErr := repo.Save("https://www.google.com", "Vsdfb1")
	secondSetErr := repo.Save("https://www.microsoft.com", "Vsdfb1")

	// Then
	assert.NoError(t, firstSetErr)
	assert.Equal(t, domain.ErrCouldNotSaveEntry, secondSetErr)
}

func TestRedisRepositoryNonExisting(t *testing.T) {
	// Given
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	repo := repository.New(client)

	_ = repo.Save("https://www.google.com", "Vsdfb1")
	_ = repo.Save("https://www.microsoft.com", "MasFgr0")

	// When
	count := repo.Count()
	entry1 := repo.Find("random")

	// Then
	assert.Equal(t, int64(2), count)
	assert.Equal(t, "", entry1)
}
