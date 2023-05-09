package main

import "fmt"

type Musica struct {
	titulo string
	artista string
	duracao int
}

func encontrarMusicas(playlists []Playlist, nomeMusica string) []Playlist {
	var resultado []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlists.musicas {
			if musica.titulo == nomeMusica {
				resultado = append(resultado,playlist)
			}
		}
	}
	return resultado
}

type Playlist