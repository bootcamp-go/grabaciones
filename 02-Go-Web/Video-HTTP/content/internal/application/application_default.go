package application

import (
	"app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Config is the configuration for the application.
type Config struct {
	// serverAddr is the address the server will bind to.
	ServerAddr string
}

// ApplicationDefault is the default implementation of the Application interface.
type ApplicationDefault struct {
	// rt is the router.
	rt *chi.Mux
	// config is the configuration for the application.
	config Config
}

// NewApplicationDefault creates a new ApplicationDefault.
func NewApplicationDefault(config Config) *ApplicationDefault {
	// default config
	defaultRouter := chi.NewRouter()
	defaultConfig := Config{
		ServerAddr: ":8080",
	}
	if config.ServerAddr == "" {
		defaultConfig.ServerAddr = config.ServerAddr
	}

	return &ApplicationDefault{
		rt:     defaultRouter,
		config: defaultConfig,
	}
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - storage
	st := make(map[int]handlers.MovieAttributes)
	// - handlers
	movieHandler := handlers.NewHandlerMovie(st, 0)

	// routes
	a.rt.Route("/movies", func(rt chi.Router) {
		// GET /movies/{id}
		rt.Get("/{id}", movieHandler.GetById())
		// POST /movies
		rt.Post("/", movieHandler.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.config.ServerAddr, a.rt)
	return
}