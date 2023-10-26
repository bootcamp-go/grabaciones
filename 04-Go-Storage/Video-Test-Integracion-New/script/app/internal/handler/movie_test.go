package handler_test

import (
	"app/internal/handler"
	"app/internal/storage"
	"app/platform/database/tester"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// DbInstance returns a new instance of sql.DB.
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

// TestHandlerMovie_GetMovies tests the GetMovies method.
func TestHandlerMovie_GetMovies(t *testing.T) {
	t.Run("succeed to get an empty list of movies", func(t *testing.T) {
		// arrange
		// - database connection
		db, err := DbInstance()
		require.NoError(t, err)
		defer db.Close()

		// - database tester: setup
		dbTester := tester.NewMySQLTester(db, "db_test_handler_movie_get_movies_empty")
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
		st := storage.NewStorageMovieMySQL(db)
		
		// - handler
		hd := handler.NewHandlerMovie(st)
		hdFunc := hd.GetMovies()

		// act
		req := http.Request{}
		res := httptest.NewRecorder()
		hdFunc(res, &req)

		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, `{"data":[]}`, res.Body.String())
		require.Equal(t, http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}, res.Header())
	})

	t.Run("succeed to get a list of movies", func(t *testing.T) {
		// arrange
		// - database connection
		db, err := DbInstance()
		require.NoError(t, err)
		defer db.Close()

		// - database tester: setup
		dbTester := tester.NewMySQLTester(db, "db_test_handler_movie_get_movies")
		defer dbTester.TearDown()

		err = dbTester.SetUp(
			`CREATE TABLE movies (
				id INT NOT NULL AUTO_INCREMENT,
				title VARCHAR(255) NOT NULL,
				year INT NOT NULL,
				director VARCHAR(255) NOT NULL,
				PRIMARY KEY (id)
			)`,
			`INSERT INTO movies (title, year, director) VALUES ('The Godfather', 1972, 'Francis Ford Coppola')`,
			`INSERT INTO movies (title, year, director) VALUES ('The Godfather: Part II', 1974, 'Francis Ford Coppola')`,
		)
		require.NoError(t, err)
		
		// - storage
		st := storage.NewStorageMovieMySQL(db)

		// - handler
		hd := handler.NewHandlerMovie(st)
		hdFunc := hd.GetMovies()

		// act
		req := http.Request{}
		res := httptest.NewRecorder()
		hdFunc(res, &req)

		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, `{"data":[{"id":1,"title":"The Godfather","year":1972,"director":"Francis Ford Coppola"},{"id":2,"title":"The Godfather: Part II","year":1974,"director":"Francis Ford Coppola"}]}`, res.Body.String())
		require.Equal(t, http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}, res.Header())
	})
}

// TestHandlerMovie_SaveMovie tests the SaveMovie method.
func TestHandlerMovie_SaveMovie(t *testing.T) {
	t.Run("succeed to save a movie", func(t *testing.T) {
		// arrange
		// - database connection
		db, err := DbInstance()
		require.NoError(t, err)
		defer db.Close()

		// - database tester: setup
		dbTester := tester.NewMySQLTester(db, "db_test_handler_movie_save_movie")
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
		st := storage.NewStorageMovieMySQL(db)

		// - handler
		hd := handler.NewHandlerMovie(st)
		hdFunc := hd.SaveMovie()

		// act
		req := http.Request{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader(
				`{"title":"The Godfather","year":1972,"director":"Francis Ford Coppola"}`,
			)),
		}
		res := httptest.NewRecorder()
		hdFunc(res, &req)

		// assert
		require.Equal(t, http.StatusCreated, res.Code)
		require.JSONEq(t, `{"data":{"id":1,"title":"The Godfather","year":1972,"director":"Francis Ford Coppola"}}`, res.Body.String())
		require.Equal(t, http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}, res.Header())
	})
}