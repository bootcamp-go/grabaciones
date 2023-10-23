Video - Linters (ppt link: https://docs.google.com/presentation/d/1HxJQTeUmeqKDnYFdxYftcBrI9BZvL9831xfKmNmeUQE/edit?usp=sharing)

---

## Slide 1 - Intro
Muy buenas a todos

---

## Slide 2 - Temario
```slide
- ¿Qué es un linter?
- Staticcheck
- Live-Coding
- Conclusión
```

En este video hablaremos de los linters, que son y para que sirven. Tambien veremos staticcheck, uno de los linters de golang, y como utilizarlo en nuestros proyectos. Luego veremos un live-coding, donde haremos uso de staticcheck en un proyecto real. Finalmente cerraremos el video con una breve conclusión de lo aprendido.

---

## Slide 3 - ¿Qué es un linter?
```slide
¿Qué es un linter?
Un linter es una herramienta que analiza el código fuente de un programa o script con el fin de detectar posibles errores y problemas

- errores de sintaxis
- errores de estilo / convencion de codigo
- potenciales problemas lógicos
- optimizacion
```

¿Qué es un linter?

Un linter es una herramienta que analiza el código fuente de un programa o script con el fin de detectar posibles errores y problemas.
Entre ellos tenemos:
- errores de sintaxis: estos pueden ser la falta de cierre de parentesis, comillas, así como problemas de indentación.
- errores de estilo / convencion de codigo: Estos errores se refieren a la forma en que se escribe y formatea el código. Incluimos casos de variables que no se usan, nombre de variables que no siguen la convención, entre otros.
- potenciales problemas lógicos: Problemas ajenos a la syntaxis, que pueden llevar a un comportamiento no deseado. Incluimos casos de variables que no se inicializan, o secciones de código que no llegan a ejecutarse por el diseño del programa.
- optimizaciones: algunos linters tienen la capacidad de sugerir una forma más eficiente de escribir el código para mejorar la performance.

En este caso usaremos un linter estático, que analiza el código sin necesidad de compilarlo o ejecutarlo.

---

## Slide 4 - Staticcheck
```slide
Staticcheck

Staticcheck es uno de los linters más populares para Go.

Instalación
- Link: https://staticcheck.io/docs/getting-started/#installation
# bash 01

Uso
# bash 02
```

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```

```bash
staticcheck . # chequea el paquete actual
staticcheck ./... # chequea todos los paquetes del proyecto
staticcheck ./<filename>.go # chequea un archivo en particular
```

Staticcheck

Staticcheck es uno de los linters más populares para Go. Este es el que haremos uso en este video. Utiliza análisis estático para encontrar bugs, problemas de rendimiento y simplificaciones, así como reglas de estilo de código. Es rapido y preciso, tambien utilizado durante continuos integration.

Cabe mencionar que go tiene un linter nativo llamado golint, asi como otras herramientas de analisis, como go vet, go fmt, goimports, aunque nos enfocaremos en staticcheck que adopta mas de 150 tipos de checks que nos serán de utilidad.

¿Como lo instalamos?
Pueden acceder al link que se encuentra en pantalla, donde indicará distintas formas de instalarlo, en nuestro caso utilizaremos el comando go install, el cual instalará el linter en los binarios de go, y nos permitirá utilizarlo desde cualquier proyecto.

¿Como lo usamos?
- Code:
    - staticcheck . # chequea el paquete actual
    - staticcheck ./... # chequea todos los paquetes del proyecto
    - staticcheck ./<filename>.go # chequea un archivo en particular

---

## Slide 5 - Conclusión
```slide
Conclusiones

Aprendimos que son los linters y la importancia de usarlos en nuestros proyectos.

Los linters son herramientas valiosas que ayudan a los programadores a mantener su código limpio, coherente y libre de errores.

Implementamos el linter staticcheck de golang.
```

Conclusión

Aprendimos que son los linters y la importancia de usarlos en nuestros proyectos.

Los linters son herramientas valiosas que ayudan a los programadores a mantener su código limpio, coherente y libre de errores, pues facilitan su lectura, mantenimiento y colaboración con otros desarrolladores.

Tambien implementamos el linter staticcheck de golang, viendo como nos ayuda a detectar errores y problemas en nuestro código.

---

## Slide 6 - Cierre

Espero que el video les haya sido de utilidad, un saludo y hasta la próxima.