## Presentacion del Proyecto
Pasemos a ver un ejemplo en código, tenemos un archivo main.go y podemos ver un struct Calculator que representa una calculadora con métodos que nos permiten ejecutar operaciones matemáticas básicas, entre ellas:
- Add: suma dos números
- Subtract: resta dos números
- Divide: divide dos números
- SquareRoot: calcula la raíz cuadrada de un número


En Divide vemos que si el numerador es 0, devuelve un error, el cual esta declarado como "DivisionError"

En SquareRoot vemos que si el número es negativo, devuelve un error, el cual esta declarado como "ErrCalculatingSquareRoot" ya que no se puede calcular la raíz cuadrada de un número negativo como número real.

De momento vemos que el programa no tiene errores de sintaxis, ni de compilación, por lo que podemos ejecutarlo y ver que funciona correctamente.

Aunque si hay detalles de mejora. Ejecutemos el linter para ver que nos dice.

> ejecutar linter
> staticcheck ./main.go

Notamos que saltan una serie de warnings:
1. error var DivisionError should have name of the form ErrFoo (ST1012): esto indica que el nombre de la variable debería ser ErrFoo siguiendo la convención de Go.

> cambiar nombre de variable a ErrDivisionZero

2. error strings should not be capitalized (ST1005): esto indica que el mensaje de error no debería estar capitalizado.

3. error strings should not end with punctuation or newlines (ST1005): por ultimo, este indica que el mensaje de error no debería terminar con un punto o un salto de linea.

> cambiar mensaje de error a "square root of negative number"