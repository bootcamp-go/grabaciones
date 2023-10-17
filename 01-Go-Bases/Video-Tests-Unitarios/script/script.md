Video Unit-Testing
- ppt link: https://docs.google.com/presentation/d/1GpTMysJCIdoYCvwXZJgyHNvXHJvSvf-adBXg039e8Xo/edit?usp=sharing
- doc link: https://docs.google.com/document/d/1zSvazIkXaInKF63RsRBC3poEg1ENxkOXOmX0QqYObYc/edit#heading=h.g3v86fqyz7f2

_______________________________________________________________
## 1a. PPT - Intro
Muy buenas a todos. En este video hablaremos de los siguientes temas

## 1b. PPT - Temario
- ¿Qué es un unit testing?
- ¿Qué características debe tener nuestros tests?
- ¿Qué forma deben tener nuestros tests?
- ¿Como escribimos nuestros tests?
- Live-Coding

```ppt
- ¿Qué es un unit testing?
- ¿Qué características debe tener nuestros tests?
- ¿Qué forma deben tener nuestros tests?
- ¿Como escribimos nuestros tests?
```

_______________________________________________________________
## 2a. PPT - ¿Qué es un unit testing?
El unit testing (o test unitario) es la práctica de escribir aquellas herramientas que nos van a permitir probar funcionalidades de nuestro código, contemplándose como partes aisladas que componen un todo. 
¿Qué queremos decir con esto? Al momento de desarrollar, estamos implementando componentes de un sistema, que luego interactuarán entre sí una vez ensamblado el producto. Pero para llegar a esa integración primero debemos asegurarnos que todas las partes funcionen correctamente, sin comportamientos extraños o impredecibles, y siendo capaces de contener posibles errores que puedan llegar a surgir durante su uso.
Es por esto, que el test unitario es una de las herramientas principales para construir un buen software.

## 2b. PPT - ¿Como debemos pensar nuestros tests?
En primer lugar, lo que tenemos que tener en claro es qué es lo que queremos probar. Parece una obviedad o algo muy evidente, pero es importante definir con rigurosidad cual va a ser el alcance de nuestro test.
Llevemos esto a un caso menos abstracto que el software:
Tenemos en nuestra casa un jardín con rosales. Y como son plantas delicadas que requieren un riego adecuado, queremos instalar un sistema de riego automático. Se va a componer de 3 piezas:
Una manguera con pequeños agujeros que dosifique el agua a nuestros rosales
Una válvula que permita el paso del agua hacia esta manguera
Un interruptor automático, que cada cierto tiempo abrirá y cerrará la válvula.

En una primera instancia vamos a querer probar que nuestra manguera dosificadora cumpla su función. Es decir, si la llenamos de agua, vaya liberando pequeñas cantidades hasta que se vacíe. ¿Y si la suciedad del piso o particulas del agua tapan nuestros dosificadores? Esa es otra prueba que queremos hacer, pero no la haremos junto con la de funcionamiento correcto. Son escenarios diferentes, para los cuales esperamos diferentes comportamientos.
¿Qué podemos observar de este ejemplo?
En primer lugar, y lo más importante, es que probamos un único componente hasta el momento: la manguera dosificadora. Pero sobre esta pieza, analizamos cuáles fueron los posibles escenarios que pueden darse, y realizamos distintas pruebas haciendo foco en cada uno de ellos. Un escenario -> una prueba

```ppt
¿Qué es un unit testing?
El unit testing es la práctica de escribir pruebas que evalúan componentes aislados de código para garantizar que funcionen correctamente antes de integrarlos en un sistema más grande.

¿Cómo debemos pensar en nuestras pruebas?
Es esencial definir claramente el alcance de las pruebas y considerar diferentes escenarios de prueba para cada componente.
```

_______________________________________________________________
## 3. PPT - ¿Que características debe tener nuestros tests?
Los tests, deben seguir un principio denominado “FIRST”, que es una sigla de las palabras
Fast (Rápido), Isolated o Independent (Aislados, independientes), Repeatable (Repetibles), Self-validating (Auto validantes) y Thorough (Minuciosos).
¿Qué quiere decir este conjunto de conceptos? 
Que los tests deben ser 
Rápidos: Si decidimos ejecutar los tests en cualquier momento del desarrollo, no deben ser una tarea de ejecución larga, por más que sean cientos de tests.
Aislados: cada test prepara su entorno, y no debe depender de algún estado externo, o afectar el entorno de otros tests
Repetibles: Debemos poder repetir nuestros tests la cantidad de veces que queramos y deben ser determinísticos. Es decir, no deben cambiar su resultado dependiendo de cuantas veces lo hayamos ejecutado o el entorno donde lo hicimos.
Auto validantes: No debemos corroborar manualmente si un test fue exitoso o falló, sino que el mismo test debe ser capaz de visualizar eso.
Minuciosos: Los tests deben probar todos los casos de éxito, probar inputs erróneos, casos borde, posibles errores irrecuperables, intentando así cubrir todos los casos de uso.

```ppt
¿Qué características deben tener nuestras pruebas?
Las pruebas deben seguir el principio "FIRST":
- Fast (Rápidas)
- Isolated o Independent (Aisladas, independientes)
- Repeatable (Repetibles)
- Self-validating (Auto validantes)
- Thorough (Minuciosas)
```

_______________________________________________________________
## 4. PPT - ¿Qué estructura deben tener nuestros tests?
Es una buena práctica ordenar nuestros tests en tres partes:
Preparación
Acción
Comprobación
Esto también es conocido como “las tres A”, por Arrange, Act y Assert

En la preparación, declararemos todas las variables, resultados esperados, seteo de variables de entorno, y dependencias que tenga nuestra pieza de código a probar.
En la acción, ejecutaremos esta pieza, y obtendremos el resultado de esta. Por si se están preguntando, en futuras clases veremos cómo comprobar la ejecución de aquellas funciones que no devuelvan valores, pero que sí realicen ciertos procesos sobre otros recursos.
Y finalmente en la comprobación, validaremos que el resultado obtenido es igual al que esperábamos

```ppt
¿Qué forma deben tener nuestras pruebas?
Las pruebas deben organizarse en tres partes: Preparación, Acción y Comprobación (conocidas como "las tres A": Arrange, Act y Assert).

- Arrange: En la preparación, se establecen las condiciones iniciales.
- Act: En la acción, se ejecuta el código bajo prueba.
- Assert: En la comprobación, se verifica si los resultados son los esperados.
```

_______________________________________________________________
## 5. PPT - ¿Como escribimos nuestros tests?
Go nos provee dentro de los paquetes nativos del lenguaje, el package “testing”. Este paquete, nos provee soporte para tests automatizados, por ejemplo, unit tests.

Para escribir un nuevo conjunto de tests, o “test suite”, debemos crear un archivo dentro de nuestro package, que termine con el sufijo “_test.go”, donde escribiremos funciones que comiencen con la palabra Test, junto con un nombre descriptivo de lo que vamos a probar.
Estas funciones deben recibir como argumento, un puntero a una instancia de la estructura T del paquete testing. Esta referencia es usada para, por ejemplo, indicar que un test falló (recordemos, auto validante), y es “inyectada” por la herramienta de ejecución de pruebas que nos provee Go: go test

No es menor aclarar, que el package al cual deben pertenecer los archivos _test.go, es al package _test. ¿Por qué? Porque queremos probar nuestro paquete desde la perspectiva de test de caja negra. Es decir, solo aquellos miembros expuestos. De esta forma, todas aquellas funciones no expuestas, son probadas indirectamente por aquellas públicas que las utilizan.

¿Qué ocurre si nuestra pieza de código presenta un panic en su ejecución?
Si el código que estamos probando “paniquea”, nuestro test falla inmediatamente

```ppt
¿Cómo escribimos nuestras pruebas en Go?
Go proporciona el paquete "testing" para escribir pruebas unitarias.

- Creamos un archivo con el sufijo "_test.go" dentro del mismo paquete que se está probando.
<imagen directorio>
- El nombre del package debe terminar con "_test" en caso de querer probar un paquete como caja negra.
- Las funciones de prueba deben comenzar con la palabra "Test" y recibir un puntero a una instancia de la estructura "T" del paquete testing.
<imagen funcion>
```

```go
package calculator_test

func TestCalculator(t *testing.T) {
    // Arrange
    // Act
    // Assert
}
```

_______________________________________________________________
## 6. PPT - ¿Cómo realizamos nuestra comprobación de resultados?
Como mencionamos previamente, nuestros tests reciben como argumento una variable de tipo puntero a “testing.T”. Ese argumento, nos servirá para marcar que un test falló, a través de la invocación del método Error o Errorf (teniendo como única diferencia, poder pasarle un string template al segundo, sin tener que nosotros construir el mensaje de error). 
Si alguno de estos métodos no se ejecuta dentro de nuestro test, el resultado del test será exitoso. Es decir, no debemos indicar que el test pasó, sino que el test falló.

La llamada a cualquiera de estos dos métodos, no detiene la ejecución del test. Por lo que si se invoca, todas las comprobaciones restantes, se ejecutarán. Si recordamos los principios FIRST, los tests deben ser rápidos. Por lo que si una validación falló, es innecesario continuar con el resto, haciendo que sean más lentos. (Recordemos que al momento de desplegar nuestro sistema, existe algún sistema de CI/CD que se encargará de ejecutar las pruebas, compilar nuestro código, y desplegarlo. Y todo tiempo de cálculo desperdiciado, es derrochar dinero)

El argumento *testing.T, posee también un método Fatal y Fatalf, que no solo indican que el test falló, sino que inmediatamente detiene la ejecución de la prueba.

```ppt
¿Cómo realizamos la comprobación de resultados?
Las pruebas en Go utilizan los metodos de testing.T Error y Errorf para indicar que una prueba falló. La llamada a estos métodos no detiene la ejecución del test. También se mencionan métodos "Fatal" que detienen la ejecución en caso de fallo.
```

```go
package calculator_test

func TestCalculator(t *testing.T) {
    // Arrange
    // Act
    // Assert
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
```

_______________________________________________________________
## 7. PPT - ¿Podemos simplificar el proceso de assertion?
El paquete “testify”, es un paquete que extiende las funcionalidades del paquete nativo “testing”, permitiendonos realizar comprobaciones de manera más compacta y con mayor nivel de precisión, ya que, por ejemplo, nos simplifica comprobar si un error es igual a un error conocido, o instancia de un tipo de error (recuerdan Errors.As()?)
Además, nos proporciona el paquete “require”, que detendrá la ejecución del test por nosotros.

```ppt
¿Podemos simplificar el proceso de aserciones?
El paquete "testify" extiende las funcionalidades del paquete "testing" y permite realizar comprobaciones de manera más compacta y precisa. También proporciona el paquete "require" que detiene la ejecución del test en caso de fallo.
```

```go
package calculator_test

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestCalculator(t *testing.T) {
    // Arrange
    // Act
    // Assert
    require.Equal(t, expected, result)
}
```

_______________________________________________________________
### 8. PPT - Conclusiones
En este documento, hemos aprendido acerca de la importancia de las pruebas unitarias en el desarrollo de software. Hemos explorado varios aspectos clave relacionados con las pruebas unitarias:
¿Qué es un unit testing?
El unit testing es la práctica de escribir pruebas que evalúan componentes aislados de código para garantizar que funcionen correctamente antes de integrarlos en un sistema más grande.

¿Cómo debemos pensar en nuestras pruebas?
Es esencial definir claramente el alcance de las pruebas y considerar diferentes escenarios de prueba para cada componente. Cada escenario debe asociarse con una prueba específica.

¿Qué características deben tener nuestras pruebas?
Las pruebas deben seguir el principio "FIRST," que significa que deben ser Rápidas, Aisladas, Repetibles, Auto validantes y Minuciosas.

¿Qué forma deben tener nuestras pruebas?
Las pruebas deben organizarse en tres partes: Preparación, Acción y Comprobación (conocidas como "las tres A": Arrange, Act y Assert). En la preparación, se establecen las condiciones iniciales; en la acción, se ejecuta el código bajo prueba; y en la comprobación, se verifica si los resultados son los esperados.

¿Cómo escribimos nuestras pruebas en Go?
Go proporciona el paquete "testing" para escribir pruebas unitarias. Las funciones de prueba deben comenzar con la palabra "Test" y recibir un puntero a una instancia de la estructura "T" del paquete testing. Los archivos de pruebas deben pertenecer al mismo paquete que se está probando.

¿Cómo realizamos la comprobación de resultados?
Las pruebas en Go utilizan métodos como Error y Errorf para indicar que una prueba falló. La llamada a estos métodos no detiene la ejecución del test. También se mencionan métodos "Fatal" que detienen la ejecución en caso de fallo.

¿Podemos simplificar el proceso de aserciones?
El paquete "testify" extiende las funcionalidades del paquete "testing" y permite realizar comprobaciones de manera más compacta y precisa. También proporciona el paquete "require" que detiene la ejecución del test en caso de fallo.
En resumen, hemos aprendido la importancia de las pruebas unitarias, cómo diseñar pruebas efectivas y algunos detalles sobre cómo escribirlas en Go. Estas pruebas son esenciales para garantizar la calidad y confiabilidad del software.

```ppt
- Unit Testing en Desarrollo de Software: El unit testing es una práctica fundamental en el desarrollo de software que implica probar componentes individuales de código de manera aislada para garantizar su correcto funcionamiento antes de la integración.

- Principio "FIRST" de las Pruebas: Las pruebas deben ser rápidas, aisladas, repetibles, auto validantes y minuciosas para ser efectivas en la identificación de errores y comportamientos inesperados.

- Estructura de las Pruebas en Go: En Go, las pruebas se escriben utilizando el paquete "testing." Las pruebas deben comenzar con la palabra "Test" y se organizan en tres partes: preparación, acción y comprobación.

- Simplificación con el Paquete "testify": El paquete "testify" extiende las funcionalidades de las pruebas en Go, permitiendo realizar comprobaciones más compactas y precisas, además de proporcionar 
```