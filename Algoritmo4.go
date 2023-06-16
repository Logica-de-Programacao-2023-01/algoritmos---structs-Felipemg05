package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo   string
	Artista  string
	Duracao  time.Duration
}

type Playlist struct {
	Nome     string
	Musicas  []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Nome da Playlist: %s\n", p.Nome)
	totalDuracao := time.Duration(0)
	for _, musica := range p.Musicas {
		fmt.Printf("Título: %s, Artista: %s, Duração: %s\n", musica.Titulo, musica.Artista, musica.Duracao.String())
		totalDuracao += musica.Duracao
	}
	fmt.Printf("Tempo total da Playlist: %s\n", totalDuracao.String())
}

func main() {
	musicas := []Musica{
		{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute + 30 * time.Second},
		{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute + 15 * time.Second},
		{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2 * time.Minute + 45 * time.Second},
	}

	p := Playlist{
		Nome:    "Minha Playlist",
		Musicas: musicas,
	}

	imprimirPlaylist(p)
}
