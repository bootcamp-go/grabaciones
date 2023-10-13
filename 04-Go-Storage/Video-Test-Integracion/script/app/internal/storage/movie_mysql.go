package storage

import (
	"app/internal"
	"database/sql"
	"errors"
)

// NewStorageMovieMySQL returns a new instance of StorageMovieMySQL.
func NewStorageMovieMySQL(db *sql.DB) *StorageMovieMySQL {
	return &StorageMovieMySQL{db: db}
}

// StorageMovieMySQL is the implementation of StorageMovie for MySQL.
type StorageMovieMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// GetMovies returns a list of movies.
func (s *StorageMovieMySQL) GetMovies() (m []internal.Movie, err error) {
	m = make([]internal.Movie, 0)
	query := `SELECT id, title, year, director FROM movies;`
	
	// execute the query
	rows, err := s.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var movie internal.Movie
		err = rows.Scan(&movie.Id, &movie.Attributes.Title, &movie.Attributes.Year, &movie.Attributes.Director)
		if err != nil {
			return
		}
		m = append(m, movie)
	}

	// check for errors during iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}

// SaveMovie saves a movie
func (s *StorageMovieMySQL) SaveMovie(m *internal.Movie) (err error) {
	query := `INSERT INTO movies (title, year, director) VALUES (?, ?, ?);`

	// prepare the statement
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec(m.Attributes.Title, m.Attributes.Year, m.Attributes.Director)
	if err != nil {
		return
	}

	// check rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = errors.New("rows affected != 1")
		return
	}

	// get the last inserted id
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return
	}

	// set the id
	(*m).Id = int(lastInsertId)

	return
}