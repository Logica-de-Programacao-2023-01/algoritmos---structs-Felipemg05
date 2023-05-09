package main

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func increment(f Funcionario, percentual float64) Funcionario {
	f.salario = f.salario + (f.salario * percentual)
	return f
}
func decrement(f Funcionario, percentual float64) Funcionario {
	f.salario = f.salario - (f.salario * percentual)
	return f
}

func calcularTempoServico(f Funcionario) int {
	return f.idade - 18
}
