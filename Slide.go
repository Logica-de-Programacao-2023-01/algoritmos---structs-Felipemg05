package main

import "fmt"

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Area(r Retangulo) float64 {
	return r.Altura * r.Largura
}

func main() {
	var r Retangulo
	fmt.Println("Informe a largura: ")
	fmt.Scan(&Largura)
	fmt.Println("Informe a largura: ")
	fmt.Scan(&Altura)
	r := Area(r)
	fmt.Println(r)
}
