# Video - Packages

---

## Slide 01 - Introduction

¡Buenas! ¿Como estan?

En este video vamos a aprender sobre paquetes y módulos en Go. ¿Estás listo? ¡Comencemos!

En ocasiones, nos encontramos con código tan extenso en un solo archivo que se vuelve ilegible y difícil de reutilizar. ¡No te preocupes, a todos nos ha sucedido! Para evitar esto, es esencial aprender a encapsular y modularizar nuestro código. En este sentido, en el lenguaje de programación Go, tenemos dos conceptos clave: paquetes (packages) y módulos (modules). Acompáñame en este viaje para descubrirlos.

---

## Slide 02 - Packages

Antes de adentrarnos, definamos estos conceptos. Un paquete (package) en Go se compone de archivos fuente que residen en un mismo directorio. Un paquete puede incluir funcionalidades adicionales que mejoran la sofisticación de tus programas.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

---

## Slide 03 - Packages

Algunos packages que podemos encontrar
- main: es el punto de entrada de una aplicación Go. Debe contener una función main, que es la función que se ejecutará cuando se inicie la aplicación.
- nativos: Go proporciona una serie de paquetes incorporados, como el paquete math, que contiene funciones matemáticas útiles, como Sqrt. Puedes importar estos paquetes en tu programa para utilizar sus funciones.
- propios: Go nos permite crear paqutes propios, que pueden contener funciones, variables, constantes, etc.

Tanto los packages nativos como los propios se pueden importar en otros packages.

Por convención los nombres de los packages son en minúsculas y sin espacios. Si el nombre del package es compuesto, se recomienda utilizar el guión bajo para separar las palabras. Para crear uno se debe crear una carpeta por convención con el mismo nombre del package que contenga los archivos. Todos los archivos dentro de la carpeta deben pertenecer al mismo package.

En el siguiente ejemplo, se crea un package llamado calculator que contiene una función Sum que recibe dos números enteros y devuelve la suma de los mismos.

```go
```dir
./calculator
    calculator.go
```

```go
package calculator

func Sum(a, b int) int {
    return a + b
}
```

---

## Slide 04 - Modules

Por otro lado, los módulos (modules) son una herramienta crucial para gestionar las dependencias de un proyecto.
Permiten la incorporación de diferentes versiones de una misma dependencia sin afectar la integridad de la aplicación. Esta característica se introdujo en Go a partir de la versión 1.11, junto con el administrador de dependencias llamado Go Modules (su version estable se incoporó a partir de la version 1.13)

Go Modules es una herramienta que se utiliza para administrar las dependencias en proyectos escritos en el lenguaje de programación Go. Te permite especificar y controlar las versiones de las dependencias externas y garantiza la integridad de tu proyecto al permitirte incorporar diferentes versiones de la misma dependencia sin conflictos. Nos permite gestionar tanto las dependencias de terceros como tus propios paquetes internos de manera eficiente.

Para utilizarlo, es necesario agregar un archivo go.mod en la raíz del proyecto, que contiene la lista de dependencias y sus versiones

Para inicializar un module en Go, hay que ejecutar el comando go mod init seguido del nombre del módulo. Por ejemplo, si queremos inicializar un módulo llamado "application”, ejecutamos el siguiente comando:

```bash
go mod init application
```

A un costado podemos ver el contenido del archivo generado go.mod, que contiene el nombre del módulo y la versión de Go utilizada.

```go.mod
module application

go 1.21
```

---

## Slide 05 - Diagrama

<diagrama>

Comprender estos dos conceptos nos lleva a la conclusión de que un módulo está compuesto por paquetes, y los paquetes están formados por archivos con extensión ".go". Esta estructura se ilustra en la siguiente imagen, donde la capa más externa representa los módulos. Inicializar un módulo es el primer paso en nuestro proyecto. La segunda capa corresponde a los paquetes. Conforme nuestro código crece, es una buena práctica encapsular funcionalidades en paquetes para mejorar la legibilidad, funcionalidad y reutilización del código. Finalmente, la capa más profunda la componen los archivos ".go" que contienen el código fuente de nuestro proyecto.

Esta organización es fundamental al comenzar cualquier proyecto en Go. Los paquetes (packages) son unidades de organización de código que mejoran la modularidad y la reutilización, mientras que los módulos (modules) son una herramienta para gestionar dependencias de terceros de manera controlada. El entendimiento de los conceptos de paquetes y módulos nos permite estructurar nuestros proyectos de manera eficiente. ¿Estás preparado para implementarlo?

---

## Slide 06 - Cierre

Espero que el video te haya sido útil. ¡Nos vemos en el próximo video!