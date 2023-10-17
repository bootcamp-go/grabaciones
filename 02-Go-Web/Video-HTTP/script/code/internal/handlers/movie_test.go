package handlers_test

import (
	"app/internal/handlers"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// Tests for HandlerMovie.GetById handler.
func TestHandlerMovie_GetById(t *testing.T) {
	t.Run("success - movie found - 200", func(t *testing.T) {
		// arrange
		// - storage
		st := make(map[int]handlers.MovieAttributes)
		st[1] = handlers.MovieAttributes{
			Title:    "The Shawshank Redemption",
			Year:     1994,
			Director: "Frank Darabont",
		}
		// - handler
		hd := handlers.NewHandlerMovie(st, 1)
		hdFunc := hd.GetById()

		// act
		req := &http.Request{}
		ctxChi := chi.NewRouteContext()
		ctxChi.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctxChi))

		res := httptest.NewRecorder()

		hdFunc(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"title":"The Shawshank Redemption","year":1994,"director":"Frank Darabont"}`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})

	t.Run("fail - invalid id - 400", func(t *testing.T) {
		// arrange
		// - storage
		st := make(map[int]handlers.MovieAttributes)
		// - handler
		hd := handlers.NewHandlerMovie(st, 0)
		hdFunc := hd.GetById()

		// act
		req := &http.Request{}
		ctxChi := chi.NewRouteContext()
		ctxChi.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctxChi))

		res := httptest.NewRecorder()

		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"invalid id",
		)
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})

	t.Run("fail - movie not found - 404", func(t *testing.T) {
		// arrange
		// - storage
		st := make(map[int]handlers.MovieAttributes)
		// - handler
		hd := handlers.NewHandlerMovie(st, 0)
		hdFunc := hd.GetById()

		// act
		req := &http.Request{}
		ctxChi := chi.NewRouteContext()
		ctxChi.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctxChi))

		res := httptest.NewRecorder()

		hdFunc(res, req)

		// assert
		expectedCode := http.StatusNotFound
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"movie not found",
		)
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})
}

// Tests for HandlerMovie.Create handler.
func TestHandlerMovie_Create(t *testing.T) {
	t.Run("success - movie created - 201", func(t *testing.T) {
		// arrange
		// - storage
		st := make(map[int]handlers.MovieAttributes)
		// - handler
		hd := handlers.NewHandlerMovie(st, 0)
		hdFunc := hd.Create()

		// act
		req := &http.Request{
			Body: io.NopCloser(
				strings.NewReader(`{"title":"The Shawshank Redemption","year":1994,"director":"Frank Darabont"}`),
			),
		}
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusCreated
		expectedBody := `{"id":1,"title":"The Shawshank Redemption","year":1994,"director":"Frank Darabont"}`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})

	t.Run("fail - invalid request body - 400", func(t *testing.T) {
		// arrange
		// - storage
		st := make(map[int]handlers.MovieAttributes)
		// - handler
		hd := handlers.NewHandlerMovie(st, 0)
		hdFunc := hd.Create()

		// act
		req := &http.Request{
			Body: io.NopCloser(strings.NewReader(
				`invalid request body`,
			)),
		}
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"invalid request body",
		)
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})

	t.Run("fail - movie year is required - 400", func(t *testing.T) {
		// arrange
		// - storage
		st := make(map[int]handlers.MovieAttributes)
		// - handler
		hd := handlers.NewHandlerMovie(st, 0)
		hdFunc := hd.Create()

		// act
		req := &http.Request{
			Body: io.NopCloser(
				strings.NewReader(`{"title":"The Shawshank Redemption","year":-1,"director":"Frank Darabont"}`),
			),
		}
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"year must be greater than or equal to 0",
		)
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})
}