package handlers

import (
	"app/platform/web/request"
	"app/platform/web/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// NewHandlerMovie creates a new HandlerMovie.
func NewHandlerMovie(st map[int]MovieAttributes, lastId int) *HandlerMovie {
	return &HandlerMovie{
		st: st,
		lastId: lastId,
	}
}

// MovieAttributes is the attributes of a movie.
type MovieAttributes struct {
	// Title is the title of the movie.
	Title string
	// Year is the year the movie was released.
	Year int
	// Director is the director of the movie.
	Director string
}

// HandlerMovie is an struct that return handlers for movies.
type HandlerMovie struct {
	// st is the storage of movies as a map.
	// - key: movie ID
	// - value: movie attributes
	st map[int]MovieAttributes
	// lastId is the last ID of a movie.
	lastId int
}

// MovieJSON is the JSON representation of a movie.
type MovieJSON struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// GetById returns a handler that returns a movie by ID.
func (h *HandlerMovie) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get movie from storage
		movie, ok := h.st[id]
		if !ok {
			response.Error(w, http.StatusNotFound, "movie not found")
			return
		}

		// response
		response.JSON(w, http.StatusOK, MovieJSON{
			Id:       id,
			Title:    movie.Title,
			Year:     movie.Year,
			Director: movie.Director,
		})
	}
}

type RequestBodyMovie struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// Create returns a handler that creates a movie.
func (h *HandlerMovie) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestBodyMovie
		if err := request.JSON(r, &reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		// - deserialize request body
		movie := MovieAttributes(reqBody)
		// - validate movie
		if movie.Title == "" {
			response.Error(w, http.StatusBadRequest, "title is required")
			return
		}
		if movie.Year < 0 {
			response.Error(w, http.StatusBadRequest, "year must be greater than or equal to 0")
			return
		}
		// - add movie to storage
		h.lastId++
		h.st[h.lastId] = movie

		// response
		response.JSON(w, http.StatusCreated, MovieJSON{
			Id:       h.lastId,
			Title:    movie.Title,
			Year:     movie.Year,
			Director: movie.Director,
		})
	}
}
