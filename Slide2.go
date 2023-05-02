package main

import "fmt"

type Livro struct {
	Titulo          string
	Autor           string
	Anodepublicacao int
}

func main() {
	livro := Livro{
		Titulo:          "Titulo",
		Autor:           "Autor",
		Anodepublicacao: 2020,
	}
	imprimeLivro(livro)
}

func imprimeLivro(b Livro) {
	fmt.Println("Titulo", b.Titulo)
	fmt.Println("Autor", b.Autor)
	fmt.Println("Ano de Publicacao", b.Anodepublicacao)
}
