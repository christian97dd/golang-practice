package main

import "fmt"

func main() {
	// [...] es para indicar que no sabes la longitud
	matriz := [...]int{10, 20, 30, 40, 50}

	matriz[0] = 1

	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	for index, value := range matriz {
		fmt.Println(index, value)
	}

	matrizB := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrizB)

	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sábado"}

	fmt.Println(diasSemana)

	diaSlice := diasSemana[0:5]

	// pushear al slice
	diaSlice = append(diaSlice, "fafafa", "fifif", "fasdfa")
	// saca el index 2 del array
	diaSlice = append(diaSlice[:2], diaSlice[3:]...)

	fmt.Println(diaSlice)
	fmt.Println(len(diaSlice))
	fmt.Println(cap(diaSlice))

	nombres := make([]string, 5)
	nombres[0] = "Christian"
	fmt.Println(nombres)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)

	copy(slice2, slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)
}
