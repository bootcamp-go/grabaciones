package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	// cfg
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "movies_test_db",
	}
	// register txdb driver
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func dbInstance(identifier string) (db *sql.DB, err error) {
	// open connection
	db, err = sql.Open("txdb", identifier)
	return
}

// TestHandlerMovie_SaveMovie tests the SaveMovie method.
func TestHandlerMovie_SaveMovie(t *testing.T) {
	t.Run("Success - movie saved", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := dbInstance("TestHandlerMovie_SaveMovie_Success")
		if err != nil {
			require.NoError(t, err)
		}
		defer db.Close()
		// - db: setup
		// ...
		// - repository
		rp := repository.NewRepositoryMovieMySQL(db)
		// - handler
		hd := handler.NewHandlerMovie(rp)
		hdFunc := hd.Save()

		// act
		request := &http.Request{
			Body: io.NopCloser(strings.NewReader(
				`{"id": 1, "title": "The Godfather", "year": 1972, "director": "Francis Ford Coppola"}`,
			)),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		}
		response:= httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusCreated
		expectedBody := `{"message":"movie created","data":{"id":1,"title":"The Godfather","year":1972,"director":"Francis Ford Coppola"}}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})

	t.Run("Failure - movie duplicated", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := dbInstance("TestHandlerMovie_SaveMovie_Failure_Duplicated")
		if err != nil {
			require.NoError(t, err)
		}
		defer db.Close()
		// - db: setup
		err = func() error {
			_, err := db.Exec(`INSERT INTO movies (id, title, year, director) VALUES (1, "The Godfather", 1972, "Francis Ford Coppola")`)
			return err
		}()
		require.NoError(t, err)
		// - repository
		rp := repository.NewRepositoryMovieMySQL(db)
		// - handler
		hd := handler.NewHandlerMovie(rp)
		hdFunc := hd.Save()

		// act
		request := &http.Request{
			Body: io.NopCloser(strings.NewReader(
				`{"id": 1, "title": "The Godfather", "year": 1972, "director": "Francis Ford Coppola"}`,
			)),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		}
		response:= httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusConflict
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"movie already exists"}`,
			http.StatusText(expectedCode),
		)
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}