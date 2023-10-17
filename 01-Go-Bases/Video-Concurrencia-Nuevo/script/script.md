Video - Concurrencia
- ppt link: https://docs.google.com/presentation/d/1D2s758VKVM-9jW7AjN6jJm5Qdh7swP9AFowk_EFx5FA/edit?usp=sharing
- doc link: https://docs.google.com/document/d/1c2QUd-xqw-J4tfWHobrh9A8DHAoR3kqob35CXaOWZ4E/edit

## 1. PPT - Intro
Muy buenas a todos, espero que esten muy bien. En este video hablaremos de los siguientes temas
- Concurrencia
- Go routines
- Channels
- Importancia de la concurrencia en nuestras aplicaciones

## 2. ¿Qué es la concurrencia?
Hasta el momento hemos visto como un programa está compuesto por instrucciones, una después de otra, y que se ejecutarán estrictamente en ese orden. Es decir, si la instrucción “X” precede a la instrucción “Y”, “Y” no será ejecutada hasta que “X” haya concluido.
Pero durante el diseño de nuestro sistema, podemos identificar que hay invocaciones que no son dependientes entre sí. Es decir, que “Y” no dependa de la salida de “X” para ejecutarse.

La concurrencia es la capacidad de un sistema para manejar múltiples tareas en forma superpuesta, aunque no necesariamente al mismo tiempo.

Pensemos el siguiente caso:
Estamos preparando una comida, pasta con salsa de tomate. Para su elaboración, tenemos que preparar la salsa, y hervir la pasta. Ambas son necesarias para el producto final, pero la cocción de la pasta no depende de que la salsa esté lista. Por lo tanto, podemos cocinar la pasta mientras la salsa se prepara.
Podemos ver este caso, como un ejemplo de concurrencia. Nos estamos ocupando de dos cosas en simultáneo, necesarias para el producto final. 
Pero hagamos énfasis en la palabra ocupando, y no haciendo dos cosas a la vez. Esto último, implicaría paralelismo. Ahora veremos un poco más sobre esto.


## 3. ¿Qué diferencia hay entre concurrencia y paralelismo?
Previamente dijimos que la concurrencia es ocuparse o encargarse de varias cosas en simultáneo, pero no hacer varias cosas en simultáneo. 
Si nos ponemos a pensar, actualmente las computadoras poseen varios núcleos o unidades de procesamiento, pero hasta los primeros años de la década del 2000 los procesadores tenían un único procesador, y sin embargo, eran capaces de darnos la ilusión de multitasking. Una computadora con un Pentium II nos permitía navegar por internet mientras escuchábamos un CD desde la lectora. En ese caso, el sistema operativo era quien manejaba la concurrencia, encargándose de varias cosas en simultáneo, pero no realizandolas al mismo tiempo. Luego, con la llegada de los procesadores con múltiples núcleos, se podían realizar tareas paralelamente, mejorando la performance general.

El paralelismo es una rama de la concurrencia, que implica que varias tareas se ejecuta en simultáneo, literalmente al mismo tiempo.

En lenguajes que no manejan concurrencia, podemos obtener un resultado similar delegando el manejo de esto al sistema operativo, levantando hilos, que se comunicarán entre sí con sockets, espacios de memoria compartidos, colas con mutexes, y otras herramientas que nosotros no vamos a ver, porque Go nos soluciona esto con las Goroutines y Channels


## 4. ¿Cómo maneja la concurrencia Go?
Al momento de ejecutar un programa Go, nuestra función main (punto de entrada principal de nuestro programa), no es el punto de entrada de nuestro ejecutable.
¿Cómo, qué dijimos? Refraseemos la idea. Nuestro código Go producirá un ejecutable una vez compilado. Una serie de instrucciones en binario (que podemos ver como lenguaje ensamblador) que se ejecutan secuencialmente. En lenguajes como C, C++ o Rust, lo primero que se ejecutará será nuestra función main. Pero en Go, lo primero que se ejecuta es el runtime system, que es un planificador y ejecutor (scheduler y dispatcher) de tareas, y es este runtime quien llama al main.
El runtime, es quien se encargará de manejar la concurrencia durante la ejecución del programa.


## 5. ¿Qué es una Goroutine?
Una Goroutine es un hilo (no se debe confundir con hilo del sistema operativo) que se ejecuta simultáneamente con el resto del programa.

Cualquier función puede ser llamada como una goroutine, con solo poner la palabra clave “go” por delante de la invocación a la función.
Esto pondrá en ejecución esa función (manejada por el runtime), y continuará con la ejecución de la instrucción que sigue.

Volviendo al caso de la instrucción X precediendo a Y, si a X la llamamos como

go X()

```go
func X() {
    fmt.Println("processing")
}
func Y() {
    fmt.Println("Hello World")
}
func main() {
    go X()
    Y()
}
```

Y, se ejecutará sin esperar a que X concluya.

¿Existe alguna limitación de cantidad de goroutines en ejecución?
No, no existe un límite. Podemos levantar la cantidad de goroutines que queramos. De hecho, dentro de la documentación del lenguaje, se nos invita a usar goroutines sin preocuparnos por su costo computacional, ya que no están implementadas a nivel sistema operativo. Sin embargo, no debemos olvidarnos que la ejecución de funciones no es gratuita, y que si bien el runtime de Go maneja eficientemente la concurrencia, en grandes cargas de goroutines (cientos de miles), probablemente debamos implementar alguna solución para evitar degradación de performance.


## 6. Channels
Pero, ¿de qué forma podemos obtener el resultado de una goroutine, o enviarle cosas durante su ejecución?

Para resolver la pregunta planteada, vienen al rescate los channels de Go. Un channel, es un tipo de dato que permite comunicar goroutines entre sí, a través de la escritura y la lectura de datos sobre estos

<diagrama>

____________________________________________________
¿Cómo creamos un nuevo channel?
Para crear un channel, debemos usar la función make, pasandole como argumento el tipo de datos, que para este caso será chan (para indicar canal), y luego el tipo de datos del canal.

```go
func main() {
    // channel of type int
    ch := make(chan int)
    // channel of type string
    ch2 := make(chan string)
}
```

____________________________________________________
¿Los canales son direccionales?
La declaración de un nuevo canal, no indica dirección. Es decir, podemos tanto leer de un canal como escribir en el mismo.
Sin embargo, si una función recibe como argumento un canal (o devuelve uno), podemos indicar que ese canal va a ser de lectura, o de escritura. Esto nos va a facilitar detectar posibles errores de escritura en canales que no deben ser escritos, sino leídos, y viceversa.
Para indicar dirección, basta con poner el símbolo <- delante del tipo de canal si es de lectura, o detrás del tipo si es de escritura. 

Ejemplo:
```go
func MyFunction(c chan<- string) {} // only-read channel

func MyFunction(c <-chan string) {} // only-write channel

func MyFunction(c chan string) {} // bidirectional channel
```


## 7. Importancia
La concurrencia es esencial en los servidores web para mejorar la eficiencia y la capacidad de respuesta.
Permite gestionar múltiples solicitudes de manera simultánea, lo que resulta en un rendimiento más rápido y escalabilidad.

La concurrencia es un concepto complicado que difiere de la programación secuencial.Es fundamental comprender esta distinción para aprovechar sus beneficios en el desarrollo de aplicaciones.
Implica nuevos patrones de programación y coordinación de tareas para evitar problemas como las condiciones de carrera. Estos patrones son esenciales para garantizar la integridad y el rendimiento de sistemas concurrentes.
Requiere un entendimiento profundo de go routines y canales en Go.Familiarizarse con estas construcciones específicas de Go es crucial para desarrollar aplicaciones concurrentes de manera efectiva en este lenguaje.

Casos de Uso
- Servidores: Atender múltiples solicitudes de clientes al mismo tiempo.
- Rate-Limits: Controlar el tráfico y garantizar una distribución justa de recursos.
- Event-Broadcasting: Sistema de difusión de eventos desde una fuente (servidor), en donde ante un evento, se notifica a múltiples destinatarios conectados simultáneamente.
- Cola de Tareas o Workers: Procesar tareas en segundo plano de manera eficiente.


En go, por ejemplo, el paquete http, el método ListenAndServe() de la estructura Server, maneja nuestras requests de forma concurrente, por lo que no tenemos que preocuparnos en implementar concurrencia explícitamente

## Conclusion
En resumen, en este video hemos explorado el concepto de concurrencia y su importancia en el mundo de la programación, centrándonos en el contexto de Go (Golang). Hemos aprendido que la concurrencia permite manejar múltiples tareas en forma superpuesta, lo que mejora la eficiencia y la capacidad de respuesta de nuestras aplicaciones.

Hemos diferenciado entre concurrencia y paralelismo, destacando que la concurrencia implica la ejecución de tareas de manera superpuesta, mientras que el paralelismo se refiere a la ejecución literalmente al mismo tiempo. En Go, las goroutines son fundamentales para lograr la concurrencia, ya que permiten ejecutar funciones de manera concurrente.

Además, hemos explorado el uso de canales (channels) en Go para comunicar y coordinar goroutines, lo que es esencial para la sincronización de tareas concurrentes. También hemos destacado la importancia de comprender profundamente las goroutines y los canales en Go para desarrollar aplicaciones concurrentes de manera efectiva.

Finalmente, hemos subrayado la importancia de la concurrencia en los servidores web, donde gestionar múltiples solicitudes de manera simultánea mejora el rendimiento y la escalabilidad. También hemos mencionado casos de uso en los que la concurrencia es crucial, como el control de rate-limits, la difusión de eventos y el procesamiento de tareas en segundo plano. En resumen, la concurrencia es una herramienta poderosa que permite mejorar la eficiencia y la capacidad de respuesta en una variedad de aplicaciones.