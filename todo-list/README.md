# Todo List

Lista de tareas interactiva en terminal hecha en Go.

## Cómo ejecutar

```bash
go run main.go
```

## Cómo se usa

1. Se muestra un menú con las opciones disponibles
2. Ingresás el número de la opción deseada
3. Después de cada acción se imprime la lista actualizada de tareas

## Opciones del menú

| Opción | Acción |
|--------|--------|
| 1 | Agregar tarea (nombre y descripción) |
| 2 | Marcar tarea como completada por índice |
| 3 | Editar tarea por índice |
| 4 | Eliminar tarea por índice |
| 5 | Salir |

## Ejemplo

```
Seleccione una opción:
 1. Agregar tarea
 2. Marcar tarea como completada
 3. Editar tarea
 4. Eliminar tarea
 5. Salir

Ingrese la opcion: 1
ingrese el nombre de la tarea: Estudiar Go
ingrese la descripcion de la tarea: Repasar punteros
tarea agregada correctamente

Lista de tareas:
===================================================
0: Estudiar Go - Repasar punteros - Completado: false
===================================================
```

## Structs y métodos

- **`Tarea`** - Struct con `nombre`, `desc` y `completado`
- **`ListaTareas`** - Struct que contiene un slice de `Tarea`
- **`agregarTarea()`** - Agrega una tarea al slice
- **`marcarCompletado()`** - Cambia `completado` a `true` por índice
- **`editarTarea()`** - Reemplaza una tarea por índice
- **`eliminarTarea()`** - Elimina una tarea del slice usando `append`

## Conceptos de Go utilizados

- Structs y métodos con receptor puntero (`*ListaTareas`)
- Slices: `append`, eliminación por índice
- `bufio.NewReader` para leer strings con espacios
- `fmt.Scanln()` para leer enteros
- Bucle `for` infinito con `return` para salir
- `switch` para manejar el menú
