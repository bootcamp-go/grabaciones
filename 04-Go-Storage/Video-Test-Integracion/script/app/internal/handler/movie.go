package handler

import (
	"app/internal"
	"app/internal/storage"
	"app/platform/web/request"
	"app/platform/web/response"
	"net/http"
)

// NewHandlerMovie returns a new instance of HandlerMovie.
func NewHandlerMovie(st storage.StorageMovie) *HandlerMovie {
	return &HandlerMovie{st: st}
}

// HandlerMovie is an struct with handler methods for the movie resource.
type HandlerMovie struct {
	// st is the storage implementation.
	st storage.StorageMovie
}

type MovieJSON struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// GetMovies returns a list of movies.
// - 200: returns a list of movies.
// - 500: internal server error.
func (h *HandlerMovie) GetMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		m, err := h.st.GetMovies()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "cannot get movies")
			return
		}

		// response
		// - serialize
		movies := make([]MovieJSON, 0)
		for _, v := range m {
			movies = append(movies, MovieJSON{
				Id:       v.Id,
				Title:    v.Attributes.Title,
				Year:     v.Attributes.Year,
				Director: v.Attributes.Director,
			})
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": movies,
		})
	}
}

// SaveMovie saves a movie
// - 201: movie created.
// - 400: bad request.
// - 500: internal server error.
type SaveMovieRequest struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}
func (h *HandlerMovie) SaveMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var req SaveMovieRequest
		err := request.JSON(r, &req)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request")
			return
		}

		// process
		// - deserialize
		mv := internal.Movie{
			Attributes: internal.MovieAttributes{
				Title:    req.Title,
				Year:     req.Year,
				Director: req.Director,
			},
		}
		// - save
		err = h.st.SaveMovie(&mv)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "cannot save movie")
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": MovieJSON{
				Id:       mv.Id,
				Title:    mv.Attributes.Title,
				Year:     mv.Attributes.Year,
				Director: mv.Attributes.Director,
			},
		})
	}
}
