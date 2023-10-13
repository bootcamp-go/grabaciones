package storage

import (
	"app/internal"
	"errors"
)

var (
	// ErrStorageMovieAlreadyExists is used when a movie already exists.
	ErrStorageMovieAlreadyExists = errors.New("storage-movie: movie already exists")
)

// StorageMovie is the interface that wraps the basic Movie methods.
type StorageMovie interface {
	// GetMovies returns a list of movies.
	GetMovies() (m []internal.Movie, err error)

	// SaveMovie saves a movie
	SaveMovie(m *internal.Movie) (err error)
}