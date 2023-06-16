package main

import "fmt"

type Animal struct {
	Nome   string
	Especie string
	Idade  int
	Som    string
}

func (a *Animal) ModificarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Especie)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Som: %s\n", a.Som)
}

func main() {
	animal := Animal{
		Nome:   "Fido",
		Especie: "Cachorro",
		Idade:  3,
		Som:    "Au Au",
	}

	animal.ImprimirInformacoes()

	animal.ModificarSom("Grrr")

	animal.ImprimirInformacoes()
}
