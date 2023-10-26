package application

import (
	"app/internal/handler"
	"app/internal/storage"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// NewApplicationDefault returns a new instance of ApplicationDefault.
func NewApplicationDefault(config *ConfigAppDefault) *ApplicationDefault {
	// default config
	defaultRouter := chi.NewRouter()
	defaultRouter.Use(middleware.Logger)
	defaultRouter.Use(middleware.Recoverer)
	defaultConfig := &ConfigAppDefault{
		ServerAddress: ":8080",
		ConfigMySQL: &mysql.Config{},
	}
	if config != nil {
		if config.ServerAddress != "" {
			defaultConfig.ServerAddress = config.ServerAddress
		}
		if config.ConfigMySQL != nil {
			defaultConfig.ConfigMySQL = config.ConfigMySQL
		}
	}

	return &ApplicationDefault{
		router:        defaultRouter,
		serverAddress: defaultConfig.ServerAddress,
		configMySQL:   defaultConfig.ConfigMySQL,
	}
}

// ConfigAppDefault is the application configuration.
type ConfigAppDefault struct {
	// ServerAddress is the server address.
	ServerAddress string
	// ConfigMySQL is the MySQL configuration.
	ConfigMySQL *mysql.Config
}

// New returns a new instance of Application.
type ApplicationDefault struct {
	// router is the router implementation.
	router *chi.Mux
	// serverAddress is the server address.
	serverAddress string
	// configMySQL is the MySQL configuration.
	configMySQL *mysql.Config
}

func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - database: open connection
	db, err := sql.Open("mysql", a.configMySQL.FormatDSN())
	if err != nil {
		return
	}
	defer db.Close()
	// - database: ping
	err = db.Ping()
	if err != nil {
		return
	}

	// - storage
	stMovie := storage.NewStorageMovieMySQL(db)

	// - handler
	hdMovie := handler.NewHandlerMovie(stMovie)

	// routes
	// - handler
	a.router.Route("/movies", func(r chi.Router) {
		// GET /movies
		r.Get("/", hdMovie.GetMovies())
		// POST /movies
		r.Post("/", hdMovie.SaveMovie())
	})

	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddress, a.router)
	return
}