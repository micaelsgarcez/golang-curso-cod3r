package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678909] = "Maria"
	aprovados[12345678908] = "Pedro"
	aprovados[12345678907] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678908])
	delete(aprovados, 12345678908)
	fmt.Println(aprovados[12345678908])

	// Map2
	funcsESalarios := map[string]float64{
		"José": 11325.45,
		"Maria": 325.45,
		"Pedro": 1325.45,
	}

	funcsESalarios["Rafael"] = 1239.91
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	// Map3 Aninhado
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira": 1456.78,
		},
		"J": {
			"José João": 11454.31,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}