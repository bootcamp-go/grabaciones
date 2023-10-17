## 1 - Presentar Scafolding
> - cmd
> - application
> - handlers

## 2 - Testing
Haremos un test unitario de los handlers GetById y Create donde necesitaremos un http.ResponseWriter y un http.Request

> crear el archivo handlers_test.go
```go
package handlers_test
/* el package se llama "handlers_test" con el sufijo "_test", de esta forma testearemos nuestros handlers como un test de caja negra. Pues no podremos acceder a los elementos no exportados. 

Procedemos creando nuestra funcion de test y agregamos casos de test para distintos escenarios. Aplicamos la triada AAA (Arrange, Act, Assert) para cada caso de test.
*/

func TestHandlerMovie_GetById(t *testing.T) {
    t.Run("success - movie found - 200", func(t *testing.T) {
        // Arrange
        // Act
        // Assert
    })

    t.Run("fail - invalid id - 400", func(t *testing.T) {
        // Arrange
        // Act
        // Assert
    })

    t.Run("fail - movie not found - 404", func(t *testing.T) {
        // Arrange
        // Act
        // Assert
    })
}
```

> explicar el arrange
Comenzamos con el arrange, donde inicializaremos nuestro storage, para luego poder crear nuestro handler y obtener el handler function

> explicar el act
Nuestro handler nos pide un http.ResponseWriter y un http.Request, por lo que debemos crearlos. Para el request lo crearemos con http.Request{}. Para el response, debemos crear un httptest.NewRecorder, el cual nos permitira obtener el response del handler.

Finalmente le pasamos los parametros al handler y ejecutamos el handler function.

Aunque queda una cosa importante por hacer. Nuestra funcion GetById obtiene le path param "id" a través de la librería chi. Chi aplica un pre-procesado del request entrante, en donde mapea los path params y los guarda en el context de *http.Request. Al testear nuestro handler desacoplado del pre-procesado de chi, debemos agregar manualmente al context de nuestro request el path param "id".

Debemos seguir los siguientes pasos:
- Crear una instancia del context de chi, al cual agregaremos el path param "id"
```go
ctxChi := chi.NewRouteContext()
ctxChi.URLParams.Add("id", "1")
```

- Luego debemos reemplazar nuestro request por un request con el mismo context pero agregando el context de chi. Esto se debe hacer así ya que http.Request no permite modificar el context.
```go
r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctxChi))
```

r.WithContext() nos permite crear un nuevo request con un context, el cual nosotros creamos con context.WithValue() donde le pasamos el context actual del request y le agregamos el context de chi con una key especifica de chi llamada chi.RouteCtxKey.

> explicar el assert
Finalmente, debemos validar el response del handler. Para esto, debemos validar el status code, el body response y los headers.
```go
expectedCode := http.StatusOK
expectedBody := `{"id":1,"title":"The Godfather","director":"Francis Ford Coppola","year":1972}`
expectedHeaders := map[string]string{
    "Content-Type": "application/json; charset=utf-8",
}
require.Equal(t, expectedCode, w.Code)
require.Equal(t, expectedBody, w.Body.String())
require.Equal(t, expectedHeaders, w.Header())
```

Tambien podemos hacer uso de require.JSONEq() para validar el body response, donde permite validar sin importar el orden de los campos.
```go
expectedBody := `{"id":1,"title":"The Godfather","director":"Francis Ford Coppola","year":1972}`
require.JSONEq(t, expectedBody, w.Body.String())
```

> ejecutar el test

> explicar los otros test directamente ya pre-codificados


---
Ahora pasaremos a testar nuestro handler Create

> mostrar el test y los casos pre-codificados

> explicar el arrange
Comenzamos con el arrange, donde inicializaremos nuestro storage, para luego poder crear nuestro handler y obtener el handler function

> explicar el act
Utilizamos http.Request{} al cual le tendremos que pasar un body. El Body es de tipo io.ReadCloser, es decir, un reader y closer combinados donde el reader permite leer de la fuente de datos o recurso y el closer cerrarlo. Para crearlo utilizaremos el package io y llamamos a la funcion io.NopCloser() que pide solo un io.Reader, el cual crearemos con strings.NewReader() con el fin de crear un reader de un string que en nuestro ejemplo es un json.

> explicar el assert

> ejecutar el test

> explicar los otros test directamente ya pre-codificados