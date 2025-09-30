package goarea

import "math"

// Pi é uma proporção matemática que representa a relação entre o perímetro de uma circunferência e seu diâmetro.
const Pi = 3.1416

// Circ calcula a área de uma circunferência dado o seu raio.
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a área de um retângulo dado sua base e altura.
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não exportada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
