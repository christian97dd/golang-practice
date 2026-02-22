package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

// un struct con un slice de tarea
type ListaTareas struct {
	tareas []Tarea
}

// metodo para agregar tareas
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para completar tarea
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// metodo para editar tareas
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// metodo para eliminarTarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	// instanciar lista de tareas

	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin)

	for {
		var option int
		fmt.Println("Seleccione una opción: \n",
			"1. Agregar tarea \n",
			"2. Marcar tarea como completada \n",
			"3. Editar tarea \n",
			"4. Eliminar tarea \n",
			"5. Salir\n")
		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t Tarea
			fmt.Print("ingrese el nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("ingrese la descripcion de la tarea")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Print("tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea a completar")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Print("tarea completada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea a actualizar")
			fmt.Scanln(&index)
			fmt.Print("ingrese el nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("ingrese la descripcion de la tarea")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Print("tarea actualizada correctamente")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea a eliminar")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Print("tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("opcion invalida")
		}
		// listar tareas de lista
		fmt.Println("Lista de tareas: ")
		fmt.Println("===================================================\n")

		for idx, tarea := range lista.tareas {
			fmt.Printf("%d: %s - %s - Completado: %t\n", idx, tarea.nombre, tarea.desc, tarea.completado)
		}
		fmt.Println("===================================================\n")
	}

}
