package storage

import (
	"app/internal"
	"app/platform/database/tester"
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func DbInstance() (db *sql.DB, err error) {
	// database configuration
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "localhost:3306",
	}

	// connect to database
	db, err = sql.Open("mysql", cfg.FormatDSN())
	return
}

// Tests for StorageMovieMySQL.GetMovies
func TestStorageMovieMySQL_GetMovies(t *testing.T) {
	t.Run("returns a empty list of movies", func(t *testing.T) {
		// arrange
		// - database connection
		db, err := DbInstance()
		require.NoError(t, err)
		defer db.Close()

		// - database tester: setup
		dbTester := tester.NewMySQLTester(db, "db_test_movie_get_movies_empty")
		defer dbTester.TearDown()

		err = dbTester.SetUp(
			`CREATE TABLE movies (
				id INT NOT NULL AUTO_INCREMENT,
				title VARCHAR(255) NOT NULL,
				year INT NOT NULL,
				director VARCHAR(255) NOT NULL,
				PRIMARY KEY (id)
			)`,
		)
		require.NoError(t, err)

		// - storage
		st := NewStorageMovieMySQL(db)
		
		// act
		m, err := st.GetMovies()

		// assert
		expectedMovies := make([]internal.Movie, 0)
		require.NoError(t, err)
		require.Equal(t, expectedMovies, m)
	})

	t.Run("returns an empty list of movies", func(t *testing.T) {
		// arrange
		// database connection
		db, err := DbInstance()
		require.NoError(t, err)
		defer db.Close()

		// database tester: setup
		dbTester := tester.NewMySQLTester(db, "db_test_movie_get_movies_full")
		defer dbTester.TearDown()

		err = dbTester.SetUp(
			`CREATE TABLE movies (
				id INT NOT NULL AUTO_INCREMENT,
				title VARCHAR(255) NOT NULL,
				year INT NOT NULL,
				director VARCHAR(255) NOT NULL,
				PRIMARY KEY (id)
			)`,
			`INSERT INTO movies (title, year, director) VALUES
				('The Shawshank Redemption', 1994, 'Frank Darabont'),
				('The Godfather', 1972, 'Francis Ford Coppola');
			`,
		)
		require.NoError(t, err)

		// storage
		st := NewStorageMovieMySQL(db)

		// act
		m, err := st.GetMovies()

		// assert
		expectedMovies := []internal.Movie{
			{Id: 1, Attributes: internal.MovieAttributes{Title: "The Shawshank Redemption", Year: 1994, Director: "Frank Darabont"}},
			{Id: 2, Attributes: internal.MovieAttributes{Title: "The Godfather", Year: 1972, Director: "Francis Ford Coppola"}},
		}
		require.NoError(t, err)
		require.Equal(t, expectedMovies, m)
	})
}

// Tests for StorageMovieMySQL.SaveMovie
func TestStorageMovieMySQL_SaveMovie(t *testing.T) {
	t.Run("saves a movie", func(t *testing.T) {
		// arrange
		// - database connection
		db, err := DbInstance()
		require.NoError(t, err)
		defer db.Close()

		// - database tester: setup
		dbTester := tester.NewMySQLTester(db, "db_test_movie_save_movie")
		defer dbTester.TearDown()

		err = dbTester.SetUp(
			`CREATE TABLE movies (
				id INT NOT NULL AUTO_INCREMENT,
				title VARCHAR(255) NOT NULL,
				year INT NOT NULL,
				director VARCHAR(255) NOT NULL,
				PRIMARY KEY (id)
			)`,
		)
		require.NoError(t, err)

		// - storage
		st := NewStorageMovieMySQL(db)
		
		// act
		mv := internal.Movie{
			Attributes: internal.MovieAttributes{
				Title: "The Shawshank Redemption",
				Year: 1994,
				Director: "Frank Darabont",
			},
		}
		err = st.SaveMovie(&mv)

		// assert
		expectedMovie := internal.Movie{
			Id: 1,
			Attributes: internal.MovieAttributes{
				Title: "The Shawshank Redemption",
				Year: 1994,
				Director: "Frank Darabont",
			},
		}
		require.NoError(t, err)
		require.Equal(t, expectedMovie, mv)
	})
}