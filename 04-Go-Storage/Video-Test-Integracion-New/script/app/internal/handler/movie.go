package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
)

// NewHandlerMovie returns a new instance of HandlerMovie.
func NewHandlerMovie(rp internal.RepositoryMovie) *HandlerMovie {
	return &HandlerMovie{rp: rp}
}

// HandlerMovie is an struct with handler methods for the movie resource.
type HandlerMovie struct {
	// rp is the repository implementation.
	rp internal.RepositoryMovie
}

type MovieJSON struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// Save saves a movie
// - 201: movie created.
// - 400: bad request.
// - 500: internal server error.
type SaveRequest struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}
func (h *HandlerMovie) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var req SaveRequest
		err := request.JSON(r, &req)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request")
			return
		}

		// process
		// - deserialize
		mv := internal.Movie{
			Id: req.Id,
			MovieAttributes: internal.MovieAttributes{
				Title:    req.Title,
				Year:     req.Year,
				Director: req.Director,
			},
		}
		// - save
		err = h.rp.Save(&mv)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRepositoryDuplicateMovie):
				response.Error(w, http.StatusConflict, "movie already exists")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "movie created",
			"data": MovieJSON{
				Id:       mv.Id,
				Title:    mv.Title,
				Year:     mv.Year,
				Director: mv.Director,
			},
		})
	}
}
