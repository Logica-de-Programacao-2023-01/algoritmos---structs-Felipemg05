package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func Area(r Retangulo) float64 {
	return r.altura * r.largura
}

func main() {
	var retangulo Retangulo
	r := Area(retangulo)
	fmt.Println(r)
}
