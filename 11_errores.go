package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ========== MANEJO DE ERRORES EN GO ==========
// Go NO usa try/catch/finally
// En su lugar, las funciones retornan un valor de error
// Por convención, el error es el ÚLTIMO valor de retorno

func main() {
	fmt.Println("=== MANEJO DE ERRORES ===\n")

	// ========== PATRÓN BÁSICO ==========
	fmt.Println("--- PATRÓN BÁSICO ---")

	// Función que puede fallar retorna (resultado, error)
	resultado, err := dividir(10, 2)
	if err != nil {
		// Manejar el error
		fmt.Println("Error:", err)
	} else {
		// Usar el resultado
		fmt.Printf("10 / 2 = %.2f\n", resultado)
	}

	// Intentar división por cero
	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error capturado:", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}

	// ========== CREAR ERRORES ==========
	fmt.Println("\n--- CREAR ERRORES ---")

	// Forma 1: errors.New()
	err1 := errors.New("este es un error simple")
	fmt.Println("Error 1:", err1)

	// Forma 2: fmt.Errorf() - permite formatear
	nombre := "archivo.txt"
	err2 := fmt.Errorf("no se pudo abrir %s", nombre)
	fmt.Println("Error 2:", err2)

	// Forma 3: Errores personalizados (ver más abajo)

	// ========== VERIFICAR ERRORES ==========
	fmt.Println("\n--- VERIFICAR ERRORES ---")

	edad, err := parsearEdad("25")
	if err != nil {
		fmt.Println("Error:", err)
		return // Salir si hay error
	}
	fmt.Printf("Edad parseada: %d\n", edad)

	// Intentar parsear valor inválido
	_, err = parsearEdad("abc")
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	// ========== IGNORAR ERRORES (NO RECOMENDADO) ==========
	fmt.Println("\n--- IGNORAR ERRORES ---")

	// Usar _ para ignorar errores (solo si estás seguro)
	valor, _ := strconv.Atoi("123") // Ignora el error
	fmt.Println("Valor (error ignorado):", valor)

	// ADVERTENCIA: Solo ignora errores si estás 100% seguro que no fallarán

	// ========== ERRORES PERSONALIZADOS ==========
	fmt.Println("\n--- ERRORES PERSONALIZADOS ---")

	// Verificar edad
	if err := verificarEdad(25); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Edad válida")
	}

	if err := verificarEdad(15); err != nil {
		fmt.Println("Error:", err)
	}

	if err := verificarEdad(150); err != nil {
		fmt.Println("Error:", err)
	}

	// ========== ERRORES CON TIPOS ==========
	fmt.Println("\n--- ERRORES CON TIPOS ---")

	// Crear error personalizado con tipo
	err = validarUsuario("", "pass123")
	if err != nil {
		// Verificar el tipo de error
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Error de validación: %s (Campo: %s)\n",
				validationErr.Message, validationErr.Field)
		} else {
			fmt.Println("Error desconocido:", err)
		}
	}

	// ========== ENVOLVER ERRORES (Wrapping) ==========
	fmt.Println("\n--- ENVOLVER ERRORES ---")

	// Desde Go 1.13: puedes envolver errores con %w
	err = procesarArchivo("config.txt")
	if err != nil {
		fmt.Println("Error completo:", err)

		// Desenvolver error para verificar el original
		if errors.Is(err, ErrArchivoNoEncontrado) {
			fmt.Println("  → Es específicamente un error de archivo no encontrado")
		}
	}

	// ========== MÚLTIPLES RETORNOS DE ERROR ==========
	fmt.Println("\n--- MÚLTIPLES OPERACIONES ---")

	resultado1, err := operacion1()
	if err != nil {
		fmt.Println("Error en operacion1:", err)
		return
	}

	resultado2, err := operacion2()
	if err != nil {
		fmt.Println("Error en operacion2:", err)
		return
	}

	fmt.Printf("Resultados: %d, %d\n", resultado1, resultado2)

	// ========== DEFER PARA LIMPIEZA ==========
	fmt.Println("\n--- DEFER PARA LIMPIEZA ---")

	// defer se ejecuta incluso si hay error
	if err := operacionConLimpieza(); err != nil {
		fmt.Println("Error en operación:", err)
	}

	// ========== PANIC Y RECOVER ==========
	fmt.Println("\n--- PANIC Y RECOVER ---")

	// panic: error irrecuperable (usar raramente)
	// recover: capturar panic (usar en defer)

	fmt.Println("Antes de función con panic")
	funcionConPanic()
	fmt.Println("Después de función con panic (no se ejecuta si no hay recover)")
}

// ========== FUNCIONES DE EJEMPLO ==========

// Función que puede fallar
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}

// Parsear edad con validación
func parsearEdad(s string) (int, error) {
	edad, err := strconv.Atoi(s)
	if err != nil {
		// Envolver el error con más contexto
		return 0, fmt.Errorf("error al parsear edad '%s': %w", s, err)
	}
	if edad < 0 || edad > 120 {
		return 0, fmt.Errorf("edad fuera de rango: %d", edad)
	}
	return edad, nil
}

// Verificar edad con múltiples condiciones
func verificarEdad(edad int) error {
	if edad < 0 {
		return errors.New("la edad no puede ser negativa")
	}
	if edad < 18 {
		return errors.New("debes ser mayor de edad")
	}
	if edad > 120 {
		return errors.New("edad no válida")
	}
	return nil // nil significa "sin error"
}

// ========== ERROR PERSONALIZADO CON TIPO ==========

// Definir tipo de error personalizado
type ValidationError struct {
	Field   string
	Message string
}

// Implementar la interfaz error
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validación falló en %s: %s", e.Field, e.Message)
}

// Usar error personalizado
func validarUsuario(username, password string) error {
	if username == "" {
		return &ValidationError{
			Field:   "username",
			Message: "no puede estar vacío",
		}
	}
	if len(password) < 8 {
		return &ValidationError{
			Field:   "password",
			Message: "debe tener al menos 8 caracteres",
		}
	}
	return nil
}

// ========== ERRORES PREDEFINIDOS ==========

var (
	ErrArchivoNoEncontrado = errors.New("archivo no encontrado")
	ErrPermisosDenegados   = errors.New("permisos denegados")
)

func procesarArchivo(nombre string) error {
	// Simular error
	err := ErrArchivoNoEncontrado

	// Envolver error con contexto adicional
	if err != nil {
		return fmt.Errorf("error al procesar %s: %w", nombre, err)
	}

	return nil
}

// ========== FUNCIONES DE EJEMPLO ADICIONALES ==========

func operacion1() (int, error) {
	// Simular operación exitosa
	return 42, nil
}

func operacion2() (int, error) {
	// Simular operación exitosa
	return 100, nil
}

func operacionConLimpieza() error {
	fmt.Println("  Iniciando operación...")

	// defer se ejecuta al final, incluso si hay error
	defer func() {
		fmt.Println("  Limpiando recursos...")
	}()

	// Simular error
	return errors.New("algo salió mal")

	// El defer se ejecuta antes de retornar
}

// ========== PANIC Y RECOVER ==========

func funcionConPanic() {
	// defer con recover captura el panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  Panic capturado:", r)
			fmt.Println("  Programa continúa...")
		}
	}()

	fmt.Println("  Antes del panic")
	panic("¡algo terrible pasó!") // Causa panic
	fmt.Println("  Esto no se ejecuta")
}

/*
CONCEPTOS CLAVE SOBRE MANEJO DE ERRORES:

1. PATRÓN BÁSICO:
   resultado, err := funcion()
   if err != nil {
       // manejar error
       return
   }
   // usar resultado

2. CONVENCIONES:
   - Error es el ÚLTIMO valor de retorno
   - nil significa "sin error"
   - Verificar SIEMPRE con if err != nil
   - Manejar o propagar el error (no ignorar)

3. CREAR ERRORES:
   - errors.New("mensaje")
   - fmt.Errorf("formato %s", variable)
   - fmt.Errorf("contexto: %w", err) // Envolver

4. INTERFAZ ERROR:
   type error interface {
       Error() string
   }
   - Cualquier tipo con método Error() string es un error

5. ERRORES PERSONALIZADOS:
   type MiError struct { ... }
   func (e *MiError) Error() string { ... }

6. ENVOLVER ERRORES (Go 1.13+):
   - fmt.Errorf("contexto: %w", err)
   - errors.Is(err, target)
   - errors.As(err, &target)

7. CUÁNDO USAR QUÉ:
   - error: Errores esperados (archivo no encontrado, red, etc.)
   - panic: Errores irrecuperables (bug en el código)
   - recover: Capturar panic (usar raramente)

8. BUENAS PRÁCTICAS:
   ✓ Siempre verificar errores
   ✓ Agregar contexto al propagar
   ✓ Retornar errores, no loguear y retornar
   ✓ Usar errores predefinidos (var ErrNotFound = ...)
   ✓ Documentar qué errores puede retornar una función

   ✗ No ignorar errores con _
   ✗ No usar panic para flujo normal
   ✗ No usar strings genéricos ("error")

9. DIFERENCIAS CON OTROS LENGUAJES:
   - JavaScript: try/catch
   - Java/C#: try/catch/finally
   - Python: try/except
   - Go: valor de retorno (más explícito)

10. VENTAJAS DEL ENFOQUE DE GO:
    - Errores explícitos (no ocultos)
    - Fuerza a pensar en errores
    - Sin overhead de excepciones
    - Control de flujo claro

EJECUTAR:
go run 11_errores.go
*/
