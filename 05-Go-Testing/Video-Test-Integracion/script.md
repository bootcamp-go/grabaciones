# Video - Test Integracion

___________________________________________________________

## PPT 01 - Intro
Muy buenas a todos, en este video aprenderemos mas sobre las pruebas de integracion.

___________________________________________________________

## PPT 03 - Test de Integracion
¿Qué es un Test de Integración?
Un test de integración es una técnica esencial en el desarrollo de software que se centra en validar la cooperación y el correcto funcionamiento conjunto de varios componentes o sistemas. En la práctica, esto implica verificar la interacción efectiva de nuestro sistema con componentes externos, tales como bases de datos, APIs, servicios de terceros, o cualquier otro elemento que opere fuera del alcance directo de nuestra aplicación.

Durante una prueba de integración, ponemos a prueba la respuesta de nuestro sistema ante una variedad de entradas (inputs). Lo crucial aquí es observar cómo el sistema procesa estas entradas y si el resultado (output) generado después de interactuar con los componentes externos cumple con las expectativas y requisitos establecidos. Esto no solo incluye la respuesta directa del sistema, sino también cómo maneja, procesa y utiliza los datos provenientes de fuentes externas.

¿Cual es su importancia?
La importancia de las pruebas de integración radica en su capacidad para detectar problemas que no son evidentes en las pruebas unitarias, que se realizan en aislamiento. Estas pruebas son fundamentales para asegurar la robustez y fiabilidad de nuestro sistema en un entorno real, donde interactúa con múltiples procesos y sistemas externos. Al realizar pruebas de integración de forma exhaustiva, no solo mejoramos la calidad y la confiabilidad del software, sino que también minimizamos riesgos y potenciales fallas que podrían surgir en la interacción con componentes externos.

___________________________________________________________

## PPT 04 - Ejemplo

Piensa en un automóvil que tiene tres partes clave: el motor, la caja de cambios y el sistema de frenos.

Pruebas Unitarias: Si solo probamos el motor, estamos haciendo una prueba unitaria. Es como encender el motor y asegurarse de que funciona, pero sin preocuparnos por las otras partes del coche.

Pruebas de Integración: Ahora, imagina que queremos asegurarnos de que todo el coche funcione bien en conjunto. Esto significa que el motor debe trabajar bien con la caja de cambios y el sistema de frenos. Por ejemplo, cuando cambiamos de marcha, el motor debe ajustar su potencia correctamente y los frenos deben responder adecuadamente. Esta prueba completa, que verifica cómo todas las partes trabajan juntas, es una prueba de integración.

En resumen, mientras que las pruebas unitarias son como verificar cada pieza del coche por separado, las pruebas de integración son como dar una vuelta con el coche para asegurarnos de que todas las piezas trabajan bien juntas.

___________________________________________________________

## PPT 05 - Casos Reales de Integracion

Ejemplos Sencillos de Pruebas de Integración

- Base de Datos Externa: Comprobar si nuestra aplicación puede guardar y recuperar datos de una base de datos que no está dentro de nuestra aplicación.

- APIs Externas: Asegurarse de que nuestra aplicación puede obtener información de APIs de terceros, como el clima o sistemas de pago.

- Sistemas de Archivos Externos: Verificar si nuestra aplicación puede leer y escribir archivos en servicios de almacenamiento en la nube, como Amazon S3 o Google Cloud Storage.

- Sistemas de Colas Externos: Comprobar que nuestra aplicación puede enviar y recibir mensajes de manera efectiva a través de sistemas de mensajería como RabbitMQ o Kafka.

___________________________________________________________

## PPT 06 - Herramientas
¿Qué herramientas podemos usar para realizar Test de Integración?
Para las pruebas de integración, es esencial usar herramientas que nos ayuden a simular y probar la interacción de nuestra aplicación con componentes externos. Aquí, los 'contenedores' juegan un papel clave.

¿Qué son los Contenedores?
Son entornos ligeros y aislados que contienen todo lo necesario para ejecutar una aplicación. Podemos configurarlos para simular el comportamiento de componentes externos, permitiéndonos probar cómo nuestra aplicación interactúa con ellos.

Herramientas de Contenedores Populares:
- Docker: Ampliamente usado para crear y gestionar contenedores.
- Podman: Una alternativa a Docker, conocido por su seguridad y facilidad de uso.
- LXC (Linux Containers): Ofrece una experiencia de contenedores más tradicional, centrada en Linux.

Estas herramientas permiten realizar pruebas de integración de manera eficiente, asegurando que nuestra aplicación funcione correctamente con sistemas externos.
___________________________________________________________

## PPT 07 - Ventajas Test Integracion
¿Por qué deberíamos realizar pruebas de integración?
Detección Temprana de Errores: Al realizar pruebas de integración, podemos identificar y corregir errores en las interacciones entre diferentes módulos del sistema antes de que avance el desarrollo. Esto ayuda a evitar problemas más complejos en etapas posteriores.

- Aseguramiento de la Calidad: Estas pruebas garantizan que los distintos componentes del sistema funcionen juntos de manera eficiente, aumentando la calidad del producto final.

- Facilita la Integración Continua: Son esenciales en entornos de integración continua, ya que permiten verificar automáticamente que los cambios recientes no afecten la funcionalidad existente.

- Mejora la Confiabilidad: Al comprobar la interacción entre distintos sistemas y componentes, las pruebas de integración incrementan la confiabilidad del software, asegurando que se comporte como se espera en un entorno de producción.

- Reducción de Costos a Largo Plazo: Aunque requieren una inversión inicial en tiempo y recursos, a largo plazo, ayudan a reducir los costos al identificar y solucionar problemas antes de que se conviertan en errores críticos.

- Validación de Interfaces y Flujos de Datos: Permiten validar la precisión de las interfaces de usuario y los flujos de datos entre diferentes módulos, asegurando una experiencia de usuario cohesiva y una operación fluida.

- Facilita la Colaboración y la Comunicación: Al requerir que distintos equipos trabajen juntos, las pruebas de integración mejoran la comunicación y colaboración entre los equipos de desarrollo, lo que resulta en un software más cohesivo.

___________________________________________________________
## PPT 10 - Cierre

Incorporar las pruebas de integración en el proceso de desarrollo no solo mejora la calidad del producto final, sino que también contribuye a un proceso de desarrollo más eficiente y colaborativo.

Espero que el video les haya sido de utilidad. Un saludo a todos y hasta la proxima.