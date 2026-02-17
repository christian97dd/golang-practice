package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// obtiene el tiempo actual declarando variables dentro del if

	if timeNow := time.Now(); timeNow.Hour() < 12 {
		fmt.Println("mañana")
	} else if timeNow.Hour() < 17 {
		fmt.Println("tarde")
	} else {
		fmt.Println("noche")
	}

	// getteo el sistema operativo
	// os := runtime.GOOS

	// lo declaro unicamente dentro del switch (fuera no va a ser accesible)
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("go run windows")
	case "linux":
		fmt.Println("go run linux")
	case "darwin":
		fmt.Println("go run macos")
	default:
		fmt.Println("go run otro os")
	}
}
