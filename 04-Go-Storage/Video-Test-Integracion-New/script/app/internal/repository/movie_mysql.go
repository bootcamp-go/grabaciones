package repository

import (
	"app/internal"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// NewRepositoryMovieMySQL returns a new instance of RepositoryMovieMySQL.
func NewRepositoryMovieMySQL(db *sql.DB) *RepositoryMovieMySQL {
	return &RepositoryMovieMySQL{db: db}
}

// RepositoryMovieMySQL is the implementation of StorageMovie for MySQL.
type RepositoryMovieMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// Save saves a movie.
func (r *RepositoryMovieMySQL) Save(m *internal.Movie) (err error) {
	query := `INSERT INTO movies (id, title, year, director) VALUES (?, ?, ?, ?);`

	// prepare the statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec(m.Id, m.Title, m.Year, m.Director)
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			switch mySQLError.Number {
			case 1062:
				err = internal.ErrRepositoryDuplicateMovie
			}
			return
		}

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

	return
}