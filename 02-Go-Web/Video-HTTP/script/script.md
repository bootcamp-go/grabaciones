Video - HTTP
- ppt link: ...

____________________________________________________________________________
## PPT 01 - Intro
Muy buenas a todos

____________________________________________________________________________
## PPT 02 - Temario
En este video vamos a ver los siguientes puntos:
- Repaso Handlers
- Test Unitario de Handlers
- Package htttest
- Live-Coding

```ppt
- Repaso Handlers
- Test Unitario de Handlers
- Package htttest
- Live-Coding
```

____________________________________________________________________________
## PPT 03 - Repaso Handlers
En el contexto de la programación web y el desarrollo de aplicaciones web, un handler es una función que se encarga de manejar una petición, evento o acción específica. En nuestro caso cuando llega una petición HTTP, el servidor web llama a la función handler que se encarga de procesarla y devolver una respuesta.

En go tenemos los handlers nativos que reciben 2 parametros: un ResponseWriter y un Request. El ResponseWriter se utiliza para escribir la respuesta que se va a devolver al cliente y el Request para obtener información de la petición.

```ppt
En el contexto de la programación web y el desarrollo de aplicaciones web, un handler es una función que se encarga de manejar una petición, evento o acción específica. 

Funcion handler: recibe un ResponseWriter y un Request.
```

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

____________________________________________________________________________
## PPT 04 - Test Unitario de Handlers
Asi como con otras funciones, para testear handlers podemos crear un test unitario. Para esto, vamos a necesitar crear una instancia para http.ResponseWriter y *http.Request.

El package http nos permite crear una instancia de Request. Podemos definir el metodo, la url y el body.
El body es un reader, es decir esta interfaz que nos permite leer de una fuente de datos en forma de stream de bytes. En este caso, como no vamos a enviar nada en el body, lo dejamos en nil.

```ppt
Para testear handlers podemos crear un test unitario, donde necesitaremos crear una instancia para http.ResponseWriter y *http.Request.

- Package http: una de sus funcionalidades nos permite crear una instancia de Request.
```

```go
request := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/"}, Body: nil}
```

Podemos hacer lo mismo con el package nativo httptest que ademas nos permite generar el ResponseWriter. httptest.NewRecorder() nos devuelve un type ResponseRecorder que implementa la interfaz ResponseWriter.

```ppt
- Package httptest: nos permite generar el ResponseWriter.
```
**Request**
```go
request := httptest.NewRequest(http.MethodGet, "/", nil)  // method, url, body
```

**ResponseWriter**
```go
response := httptest.NewRecorder() // retorna *httptest.ResponseRecorder
```

Conociendo esto, podemos pasar a un live-coding para ver como testear un handler.

____________________________________________________________________________
## PPT 05 - Live Coding
./code/

____________________________________________________________________________
## PPT 06 - Conclusiones
En resumen, en esta presentación hemos explorado el concepto de handlers en el contexto de peticiones HTTP y cómo se utilizan en Go. Hemos aprendido que un handler es una función que se encarga de manejar una petición, evento o acción específica, y en Go, se definen con parámetros ResponseWriter y Request. Además, hemos visto cómo realizar pruebas unitarias de handlers, creando instancias simuladas de ResponseWriter y Request para asegurarnos de que nuestros manejadores funcionan correctamente. En el siguiente segmento, se llevará a cabo un live coding para poner en práctica lo aprendido.

```ppt
- Handlers HTTP en Go: Concepto esencial para el manejo de peticiones web.
- Estructura de un Handler: Respuesta y solicitud (ResponseWriter y Request).
- Pruebas Unitarias de Handlers: Creación de instancias simuladas de ResponseWriter y Request para pruebas.
- Live Coding: Aplicación práctica de lo aprendido.
```