# Guess Number Game

Juego de adivinar un número aleatorio en Go.

## Cómo ejecutar

```bash
go run main.go
```

## Cómo se juega

1. El programa genera un número aleatorio entre **0 y 99**
2. Tienes **10 intentos** para adivinarlo
3. Después de cada intento te dirá si el número es **mayor** o **menor**
4. Si lo adivinas o se acaban los intentos, puedes elegir jugar de nuevo

## Ejemplo

```
Ingrese un numero (intentos restantes: 10):  50
El numero es mayor
Ingrese un numero (intentos restantes: 9):  75
El numero es menor
Ingrese un numero (intentos restantes: 8):  62
adivinaste el numero

jugar de nuevo? (s/n):
```

## Funciones

- **`main()`** - Punto de entrada, llama a `play()`
- **`play()`** - Lógica principal del juego, genera el número y maneja los intentos
- **`playAgain()`** - Pregunta si el usuario quiere jugar otra vez

## Conceptos de Go utilizados

- `math/rand` para generar números aleatorios
- `fmt.Scanln()` para leer input del usuario
- Bucle `for` con contador de intentos
- Condicionales `if / else if`
- `switch` para manejar la respuesta de jugar de nuevo
- Recursión en `playAgain()` para reintentar si la entrada es inválida
