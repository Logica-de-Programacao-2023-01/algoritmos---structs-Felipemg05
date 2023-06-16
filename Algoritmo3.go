package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func calcularArea(t Triangulo) float64 {
	return t.Base * t.Altura / 2
}

func main() {
	t := Triangulo{
		Base:   4.5,
		Altura: 3.2,
	}
	area := calcularArea(t)
	fmt.Println("Área do triângulo:", area)
}
