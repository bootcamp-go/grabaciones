Video - Linters (ppt link: https://docs.google.com/presentation/d/1HxJQTeUmeqKDnYFdxYftcBrI9BZvL9831xfKmNmeUQE/edit?usp=sharing)

---

## Slide 1 - Intro
Muy buenas a todos, en este video hablaremos de los linters, que son y para que sirven.

---

## Slide 2 - ¿Qué es un linter?
```slide
¿Qué es un linter?

Un linter es una herramienta que analiza el código fuente de un programa o script con el fin de detectar posibles errores y problemas
- errores de sintaxis
- errores de estilo / convencion de codigo
- potenciales problemas lógicos
- optimizacion
```

```orador
¿Qué es un linter?

Un linter es una herramienta que analiza el código fuente de un programa o script con el fin de detectar posibles errores y problemas.
Entre ellos tenemos:
- errores de sintaxis: estos pueden ser la falta de cierre de parentesis, comillas, así como problemas de indentación.
- errores de estilo / convencion de codigo: Estos errores se refieren a la forma en que se escribe y formatea el código. Incluimos casos de variables que no se usan, nombre de variables que no siguen la convención, entre otros.
- potenciales problemas lógicos: Problemas ajenos a la syntaxis, que pueden llevar a un comportamiento no deseado. Incluimos casos de variables que no se inicializan, o secciones de código que no llegan a ejecutarse por el diseño del programa.
- optimizaciones: algunos linters tienen la capacidad de sugerir una forma más eficiente de escribir el código para mejorar la performance.

En este caso usaremos un linter estático, que analiza el código sin necesidad de compilarlo o ejecutarlo.
```

---

## Slide 3 - Staticcheck
```slide
Staticcheck

En nuestro caso utilizaremos el linter standard de golang, llamado staticcheck

¿Como lo instalamos?
- Link: https://staticcheck.io/docs/getting-started/#installation
- Code:
    - go install honnef.co/go/tools/cmd/staticcheck@latest

¿Como lo usamos?
- Code:
    - staticcheck . # chequea el paquete actual
    - staticcheck ./... # chequea todos los paquetes del proyecto
    - staticcheck ./<filename>.go # chequea un archivo en particular
```

```orador
Staticcheck

En nuestro caso utilizaremos el linter standard de golang, llamado staticcheck, el cual utilizando analisis estatico, nos permitira encontrar bugs, problemas de performance, así como ofrecer simplificaciones y reglas de estilo de código

¿Como lo instalamos?
Pueden acceder al link que se encuentra en pantalla, donde indicará distintas formas de instalarlo, en nuestro caso utilizaremos el comando go install, el cual instalará el linter en los binarios de go, y nos permitirá utilizarlo desde cualquier proyecto.

¿Como lo usamos?
- Code:
    - staticcheck . # chequea el paquete actual
    - staticcheck ./... # chequea todos los paquetes del proyecto
    - staticcheck ./<filename>.go # chequea un archivo en particular
```

---

## Slide 4 - Conclusión
```slide
Conclusión

Aprendimos que son los linters y la importancia de usarlos en nuestros proyectos.

Implementamos el linter standard de golang 

Los linters son herramientas valiosas que ayudan a los programadores a mantener su código limpio, coherente y libre de errores.
```

```orador
Conclusión

Aprendimos que son los linters y la importancia de usarlos en nuestros proyectos.

Los linters son herramientas valiosas que ayudan a los programadores a mantener su código limpio, coherente y libre de errores, pues facilitan su lectura, mantenimiento y colaboración con otros desarrolladores.
```