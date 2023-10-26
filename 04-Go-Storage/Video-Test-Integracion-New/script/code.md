________________________________________________________________
## Introducción Proyecto
- Mostrar el schema Movie del archivo `./docs/db/movies_db.sql`
- Mostrar capa repository: esta se encarga de interactuar con la base de datos
- Mostrar capa handler: esta se encarga de manejar las peticiones del usuario e interactuar con la capa storage
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

Recordando los tests, estos deben realizarse en base a los principios FIRST, donde nos enfocaremos en su aislamiento y repitibilidad

Previamente mencionamos que utilizaremos una base de datos compartidas junto con el package go-txdb, donde se aplica rollbacks y aislamiento.

Comenzemos creando el test de HandlerMovie para el metodo Save, un caso de success, donde se guardara una pelicula en la base de datos, y un caso
failure donde la pelicula se encuentra duplicada.
```

### Test HandlerMovie
**Test HandlerMovie.Save**
> crear archivo `./handler/movie_test.go`
> 
```go
package handler_test

// Tests for HandlerMovie.Save
func TestHandlerMovieSave(t *testing.T) {
	t.Run("Success - movie saved", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("Failure - movie duplicated", func(t *testing.T) {
		// arrange

		// act

		// assert
	})
}
```

Ahora instalaremos el package go-txdb
```bash
go get github.com/DATA-DOG/go-txdb
```

Pasamos a registrar el driver de txdb, sobre el de mysql. Para ello necesitaremos pasar el nombre del driver txdb, luego menciona el driver el cual abtraera, mysql y finalmente la conexion a la db.
En go la funcion `init` es una función especial que se utiliza para inicializar paquetes dentro de un programa. La función init se ejecuta automáticamente antes de que se inicie la función main en un programa Go. Cada paquete en Go puede tener una o varias funciones init, y todas se ejecutarán en el orden en que se importan los paquetes.
En nuestro caso se ejecutara previo a los tests para registrar el driver de txdb sobre el de mysql. Una vez hecho esto, las conexiones que abramos con sql.Open() inician una transaccion aislada y con rollback, por lo que no se guardaran los cambios en la base de datos.
```go
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
```

Pasamos con el arrange, donde primero inicializaremos la db, luego podremos aplicar una serie de set-ups en caso de ser necesario. Para este test con la db vacia es suficiente. Luego creamos el repository y finalmente el handler del cual obtendremos el handler Save.
```go
unc TestHandlerMovieSave(t *testing.T) {
	t.Run("Success - movie saved", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := sql.Open("txdb", "TestHandlerMovieSave_Success")
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
	})
}
```

Pasamos con el act, donde llamaremos al handler el cual necesita un *http.Request y un http.ResponseWriter
http.Request pide para el body un io.ReadCloser, otra interfaz. Para ello podremos usar io.NopCloser, el cual pide tambien un io.Reader, otra interfaz. Para ello podremos usar strings.NewReader, el cual pide un string
transformando la cadena de texto en un reader, del cual podremos leer como un stream de datos en bytes.
Agregamos los headers indicando que el contenido es en formato json
Luego creamos el response con httptest.NewRecorder() el cual devuelve un struct httptest.ResponseRecorder que implementa la interfaz http.ResponseWriter
finalmente ejecutamos el handler
```go
func TestHandlerMovieSave(t *testing.T) {
	t.Run("Success - movie saved", func(t *testing.T) {
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
	})
}
```

Finalmente armamos el assert. Recuerde require.JSONEq nos permite comparar json sin importar el orden de las claves.
```go
func TestHandlerMovieSave(t *testing.T) {
	t.Run("Success - movie saved", func(t *testing.T) {
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
}
```

---

**Test HandlerMovie.Save**
> mostrar el test completo e indicar las diferencias
El otro caso de test es muy similar, con unas diferencias

En el setup, insertamos en al db una pelicula con una funcion. En caso de fallar el test no podra continuar, por eso verificamos con require.NoError()

Luego instanciamos lo necesario hasta obtener el handler Save.

Enviamos el mismo request, pero como la pelicula ya existe, ahora esperamos que haya un error 409, ya que la db fallaría indicando el codigo 1062 de Duplicate entry.

Para el expectedBody como utilizamos la funcion response.Error(), la estructura del mensaje es ...
> mostrar errorResponse
por ende el expectedBody quedaría como ...
```go
	expectedBody := fmt.Sprintf(
		`{"status":"%s","message":"movie already exists"}`,
		http.StatusText(expectedCode),
	)
```