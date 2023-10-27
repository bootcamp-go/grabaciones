# Video - Test Integracion (ppt link: https://docs.google.com/presentation/d/1hfFLFl9qJElON0irBHqW2lc3uL0cyVB7K9fUKA8_dvw/edit?usp=sharing)

________________________________________________________________________________________________

## PPT 01 - Intro
Muy buenas a todos

________________________________________________________________________________________________

## PPT 02 - Temario

En este video veremos los siguientes temas:
- Repaso de Tests Unitarios
- ¿Qué es un Test de Integración?
- Herramientas para llevar a cabo un Test de Integración
- Implementacion
- Package go-txdb
- Live Coding

________________________________________________________________________________________________

## PPT 03 - Repaso Test Unitario
Antes de adentrarnos en las pruebas de integración, es crucial tener una comprensión sólida de las pruebas unitarias.

Los tests unitarios son pruebas unitarias esenciales para verificar que una unidad de código dentro de nuestro programa funcione correctamente. 

Su objetivo es confirmar que, ante diferentes entradas (inputs), el proceso produce el resultado esperado. 

Estas pruebas pueden ser de dos tipos: de caja blanca (cuando conocemos el código de la unidad probada) o de caja negra (cuando no conocemos el código subyacente).

________________________________________________________________________________________________

## PPT 04 - Test Integracion
¿Qué es un Test de Integración?
Un test de integración es una prueba enfocada en verificar la correcta interacción de nuestro sistema con componentes externos, como bases de datos o servicios externos, que representan procesos ajenos a nuestra aplicación. El proceso aplicado por estos componentes está abstraído y queda fuera de nuestro control.

Durante el test, evaluaremos si ante cierto input, el output que nuestro sistema genera después del proceso aplicado, incluyendo la interacción con la data obtenida de estos componentes externos, es el esperado.

Las pruebas de integración son cruciales para asegurar que nuestro sistema funcione sin problemas con procesos externos.

________________________________________________________________________________________________

## PPT 05 - Herramientas
¿Qué herramientas podemos usar para realizar Test de Integración?
Para realizar pruebas de integración efectivas, necesitamos utilizar las herramientas adecuadas. 

Para simular el comportamiento aplicado por estos componentes externos, podemos hacer uso de Docker o Test Containers.

- Docker: Es una herramienta versátil que facilita la creación, implementación y ejecución de contenedores de software.
- Test Containers: Instancias ligeras ya pre-configuradas para levantar servicios como bases de datos, message brokers, web browsers, entre otros.

Estos containers son entornos aislados y ligeros que contienen todo lo necesario para ejecutar una aplicacion.

Recordando los principios FIRST: Fast, Isolated, Repeatable, Self-Validating, Throurough.
Si cada prueba trabaja con su propia instancia que simular el comportamiento del componente externo, podemos asegurar el aislamiento de cada prueba, y que cada una de ellas sea independiente de las demás, ademas de ser repetibles generando el mismo resultado.

________________________________________________________________________________________________

## PPT 06 - ¿Como lo implementaremos?

En nuestro caso utilizaremos un approach un poco mas rígido por cuestiones de recursos, pero que nos permitirá realizar pruebas de integración efectivas.
Utilizaremos una base de datos compartida para diferentes sets de pruebas.
<!-- Esto implica un concepto importante a tener en cuenta: `shared memory`. -->


<!-- Shared memory es un recurso compartido entre diferentes procesos, en este caso, entre diferentes pruebas. Algunas aclaraciones a tener en cuenta: -->
<!-- - Mutable: en entornos concurrente, la memoria compartida mutable corre el riesgo de `data race`, es decir una condicion de concurrencia que ocurre cuando 2 o mas hilos de ejecucion acceden simultaneamente a un mismo recurso compartido, y al menos uno de ellos realiza una operacion de escritura. Esto puede generar inconsistencias en los datos. Para ello se aplicarían mecanismos de sincronizacion como mutex, limitando el acceso al recurso compartido. -->
<!--  -->
<!-- - Inmutable: en entornos concurrentes, la memoria compartida inmutable no requiere necesariamente sincronizar el acceso, ya que no se puede modificar y las copias de lectura son seguras. -->

---

<!-- Retomando nuestro caso, que trabajaremos sobre una base de datos compartida y mutable. -->
Para que nuestros test respeten isolation (aislado / independiente) y repeatable (repetible) deberan seguir los siguientes principios:
- Condiciones iniciales: cada test debería ejecutarse en las mismas condiciones iniciales, por ende nuestra base de datos compartida debería estar limpia al inicio de cada test y luego será configurada internamente por cada test.
- Aislado: Cada test debería ejecutarse en forma aislada de los cambios de otros tests.

Si los tests se llegan a ejecutar simultaneamente (go lo permite con t.Parallel()), deberiamos asegurarnos de que si un test accede a una memoria compartida, no se permita a otro test acceder, hasta que el primero haya finalizado, en casos donde alguno de ellos aplique operaciones de escritura.

De esta forma podemos asegurar que cada test trabaje con las mismas condiciones iniciales, en forma aislada y puedan ser repetibles.

________________________________________________________________________________________________

## PPT 07 - Package go-txdb

Para ello, existe un package go-txdb que nos permite trabajar con una base de datos compartida, aplicando 2 procesos importantes:
- rollback: al cerrar la conexion con la db, en nuestros casos cuando finalice cada test, se aplica un rollback de la base de datos, dejandola en el estado inicial.
- isolation: cada test trabaja en su propia transacción, aislando los cambios de otros tests.

¡Aclaracion! en DBMS, es decir en sistemas gestores de base de datos, como MySQL y Postgres el rollback no aplica en ciertos casos:
- Cambios en la estructura de la base de datos: el lenguaje DDL, entre ellos crear tablas, alterar tablas, crear indices, etc.
- Autoincrementales: si insertamos un registro con un autoincremental, al realizar el rollback, el autoincremental no se reinicia a 0, sino que continua con el ultimo valor insertado.

Esto nos limita a trabajar con una base de datos con una estructura fija (tablas pre-definidas) y sin autoincrementales.

¿Como lo instalamos?
```bash
go get github.com/DATA-DOG/go-txdb
```

________________________________________________________________________________________________

## PPT 08 - Live-Coding
./code/

________________________________________________________________________________________________

## PPT 09 - Conclusiones
Como cierre, recapitulemos lo que hemos aprendido en este video.

En este video, exploramos a fondo las pruebas de integración, fundamentales para verificar la interacción de nuestro sistema con componentes externos.

Analizamos las herramientas disponibles para llevar a cabo pruebas de integración, como Docker y Test Containers.

Finalmente vimos una implementación de test de integración con el package go-txdb. Vimos casos de success y error.

________________________________________________________________________________________________

## PPT 10 - Cierre

Espero que el video les haya sido de utilidad. Un saludo a todos y hasta la proxima.