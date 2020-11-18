package fib

// Fib calcula el 6to número de la secuencia fibonacci.
func Fib() int {
	return Recursivo(6)
}

// Ciclo func
func Ciclo(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// Recursivo func
func Recursivo(n int) int {
	if n < 2 {
		return n
	}

	return Recursivo(n-1) + Recursivo(n-2)
}
