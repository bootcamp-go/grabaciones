package internal

import "errors"

var (
	// ErrRepositoryDuplicateMovie is returned when a movie already exists.
	ErrRepositoryDuplicateMovie = errors.New("repository: movie already exists")
)

// RepositoryMovie is an interface with methods for the movie resource.
type RepositoryMovie interface {
	// Save saves a movie.
	Save(m *Movie) (err error)
}