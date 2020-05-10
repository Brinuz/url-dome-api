package repository_test

import (
	"errors"
	"testing"
	repository "url-at-minimal-api/internal/external_interfaces/repository/postgres"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostgresRepositorySave(t *testing.T) {
	// Given
	db, mock, _ := sqlmock.New()
	repo := repository.New(db)

	mock.ExpectExec("INSERT INTO urls").WithArgs("https://www.google.com", "Vsdfb1").WillReturnResult(sqlmock.NewResult(1, 1))

	// When
	err := repo.Save("https://www.google.com", "Vsdfb1")

	// Then
	expectations := mock.ExpectationsWereMet()
	assert.NoError(t, expectations)
	assert.NoError(t, err)
}

func TestPostgresRepositoryFailedSave(t *testing.T) {
	// Given
	db, mock, _ := sqlmock.New()
	repo := repository.New(db)

	mock.ExpectExec("INSERT INTO urls").
		WithArgs("https://www.google.com", "Vsdfb1").
		WillReturnError(errors.New("dummy"))

	// When
	err := repo.Save("https://www.google.com", "Vsdfb1")

	// Then
	expectations := mock.ExpectationsWereMet()
	assert.NoError(t, expectations)
	assert.Error(t, err)
}

func TestPostgresRepositoryFind(t *testing.T) {
	// Given
	db, mock, _ := sqlmock.New()
	repo := repository.New(db)

	mock.ExpectQuery("SELECT url FROM urls").
		WithArgs("Vsdfb1").
		WillReturnRows(sqlmock.NewRows([]string{"url"}).AddRow("https://www.google.com"))

	// When
	result := repo.Find("Vsdfb1")

	// Then
	expectations := mock.ExpectationsWereMet()
	assert.NoError(t, expectations)
	assert.Equal(t, "https://www.google.com", result)
}

func TestPostgresRepositoryNoResultFind(t *testing.T) {
	// Given
	db, mock, _ := sqlmock.New()
	repo := repository.New(db)

	mock.ExpectQuery("SELECT url FROM urls").
		WithArgs("Vsdfb1").
		WillReturnRows(sqlmock.NewRows([]string{"url"}))

	// When
	result := repo.Find("Vsdfb1")

	// Then
	expectations := mock.ExpectationsWereMet()
	assert.NoError(t, expectations)
	assert.Equal(t, "", result)
}

func TestPostgresRepositoryFailedFind(t *testing.T) {
	// Given
	db, mock, _ := sqlmock.New()
	repo := repository.New(db)

	mock.ExpectQuery("SELECT url FROM urls").
		WithArgs("Vsdfb1").
		WillReturnError(errors.New("dummy"))

	// When
	result := repo.Find("Vsdfb1")

	// Then
	expectations := mock.ExpectationsWereMet()
	assert.NoError(t, expectations)
	assert.Equal(t, "", result)
}
