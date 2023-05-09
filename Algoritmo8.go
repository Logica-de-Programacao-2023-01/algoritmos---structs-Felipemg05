package main

import "fmt"

type Viagem struct {
	origem string
	destino string
	data int
	preco float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	var viagemMax float64
	for _, viagem := range viagens {
		if viagem.preco > viagemMax.preco {
			viagemMax =
		}
	}
}