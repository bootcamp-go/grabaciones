________________________________________________________________
## Introducción Proyecto
- Mostrar el schema Movie
- Mostrar capa storage: esta se encarga de interactuar con la base de datos
- Mostrar capa handlers: esta se encarga de manejar las peticiones del usuario e interactuar con la capa storage
- Mostrar capa application: esta se encarga de manejar la app con su configuración enfocada en los requisitos finales.
- Mostrar capa main: esta se encarga de ejecutar la aplicación.

________________________________________________________________
## Ejecución
- Ejecutar el comando `go run main.go` para ejecutar la aplicación para mostrar que el código ejecuta correctamente.

________________________________________________________________
## Pruebas
Menciona que haremos una prueba de integración para mostrar que la aplicación funciona correctamente interactuando con la base de datos.

Indicar que el test lo realizaremos desde la capa Handlers

```
Speech: al revisar que necesitaremos para realizar nuestro test de integración, vemos que necesitamos una instancia de HandlerMovie, la
cual tiene una dependencia que es StorageMovie, por lo que necesitamos una instancia de StorageMovie. Esta ultima tiene otra dependencia, la propia
base de datos, por lo que necesitamos una instancia de la base de datos.

Recordando los tests, estos deben realizarse de forma aislada, pues cada test
que ejecutemos debería interactuar con una propia instancia de la base de datos, por lo que necesitamos una base de datos por test.

Previamente vimos herramientas como Test Containers en los que podemos levantar un servicio como una base de datos. Podríamos usar esta herramienta
para simular su comportamiento, pero en este caso y con fines prácticos, usaremos una base de datos local, junto a un package que nos permitirá
crear schemas de la db por test.

Este package creara un schema de db con el nombre que le indiquemos, luego prepararemos la misma con queries y al final del test, la eliminaremos.
```

________________________________________________________________
## Package Tester
- Mostrar package "app/platform/database/tester", el archivo "mysql.go"

________________________________________________________________
## Test HandlerMovie
**Test HandlerMovie.GetMovies**
- Crear la funcion DbInstance()
```go
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
```

- Crear la funcion TestHandlerMovie_GetMovies() y agregar test cases
```go
// TestHandlerMovie_GetMovies tests the GetMovies method.
func TestHandlerMovie_GetMovies(t *testing.T) {
    t.Run("succeed to get an empty list of movies", func(t *testing.T) {
        // arrange

        // act

        // assert
    })

    t.Run("succeed to get a list of movies", func(t *testing.T) {
        // arrange

        // act

        // assert
    })
}
```

- Comenzar con `arrange` en el caso "succeed to get an empty list of movies"
```go
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

        // ...
    })
```

- Continuar con storage y handler
```go
	t.Run("succeed to get an empty list of movies", func(t *testing.T) {
		// arrange
        // - storage
		st := storage.NewStorageMovieMySQL(db)
		
		// - handler
		hd := handler.NewHandlerMovie(st)
		hdFunc := hd.GetMovies()
    })
```

- Continuar con `act` explicando que el handler recibe como parametros un request y un response, por lo que tendremos que crear un request utilizando el package http.Request y un response. El response es una interface http.ResponseWriter. Junto al package httptest podremos crear una instancia de un type que lo implementa, con httptest.NewRecorder() que retorna httptest.ResponseRecorder{}
```go
func TestHandlerMovie_GetMovies(t *testing.T) {
	t.Run("succeed to get an empty list of movies", func(t *testing.T) {
		// act
		req := http.Request{}
		res := httptest.NewRecorder()
		hdFunc(res, &req)
    })
```

- Finalmente, `assert` verificando que el status code, el body, y los headers
```go
	t.Run("succeed to get an empty list of movies", func(t *testing.T) {
		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, `{"data":[]}`, res.Body.String())
		require.Equal(t, http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}, res.Header())
	})
```

- Continuar sin problema con "succeed to get a list of movies"

________________________________________________________________
## Test HandlerMovie.SaveMovie
Lo mismo. Cuando se llegue a la parte del request, explicar como armar el body. http.Request pide para el body un io.ReadCloser, otra interfaz. Para ello podremos usar io.NopCloser, el cual pide tambien un io.Reader, otra interfaz. Para ello podremos usar strings.NewReader, el cual pide un string
transformando la cadena de texto en un reader, del cual podremos leer como un stream de datos en bytes.
```go
    // act
    req := http.Request{
        Header: http.Header{
            "Content-Type": []string{"application/json"},
        },
        Body: io.NopCloser(strings.NewReader(
            `{"title":"The Godfather","year":1972,"director":"Francis Ford Coppola"}`,
        )),
    }
```