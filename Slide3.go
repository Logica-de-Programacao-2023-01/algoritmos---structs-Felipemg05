package main

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func main() {
	a := Aluno{
		Nome:  "Felipe",
		Idade: 18,
		Notas: []float64{1, 2, 3, 4, 5},
	}
}

func media(a Aluno) float64 {
	var sum float64
	for _, nota := range a.Notas {
		sum += nota
	}
	return sum / float64(len(a.Notas))
}