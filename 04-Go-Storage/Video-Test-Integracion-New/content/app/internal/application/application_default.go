package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigAppDefault is the application configuration.
type ConfigAppDefault struct {
	// Router is the router implementation.
	Router *chi.Mux
	// ServerAddress is the server address.
	ServerAddress string
	// ConfigMySQL is the MySQL configuration.
	ConfigMySQL *mysql.Config
}

// NewApplicationDefault returns a new instance of ApplicationDefault.
func NewApplicationDefault(config *ConfigAppDefault) *ApplicationDefault {
	// default config
	defaultConfig := &ConfigAppDefault{
		Router:        chi.NewRouter(),
		ServerAddress: ":8080",
		ConfigMySQL: &mysql.Config{},
	}
	if config != nil {
		if config.Router != nil {
			defaultConfig.Router = config.Router
		}
		if config.ServerAddress != "" {
			defaultConfig.ServerAddress = config.ServerAddress
		}
		if config.ConfigMySQL != nil {
			defaultConfig.ConfigMySQL = config.ConfigMySQL
		}
	}

	return &ApplicationDefault{
		router:        defaultConfig.Router,
		serverAddress: defaultConfig.ServerAddress,
		configMySQL:   defaultConfig.ConfigMySQL,
	}
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

// SetUp sets up the application.
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
	// - repository
	rp := repository.NewRepositoryMovieMySQL(db)
	// - handler
	hd := handler.NewHandlerMovie(rp)

	// routes
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - handler
	a.router.Route("/movies", func(r chi.Router) {
		// POST /movies
		r.Post("/", hd.Save())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddress, a.router)
	return
}