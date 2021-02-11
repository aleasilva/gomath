package calculo

import (
	"fmt"
	"math"
)

// Pi é número do Pi
const Pi = 3.141516

// Circ Calcular a área de circula
func Circ(raio float64) float64 {
	fmt.Println(raio)
	return math.Pow(raio, 2) * Pi
}
