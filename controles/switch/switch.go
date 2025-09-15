package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int: 
		return "inteiro"
	case float32, float64: 
		return "real"
	case string: 
		return "string"
	case func(): 
		return "função"
	default: 
		return "não sei"
	}
}

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch {
		case nota >= 9 && nota <= 10:
			return "A"
		case nota >= 7 && nota < 9:
			return "B"
		case nota >= 5 && nota < 7:
			return "C"
		case nota >= 3 && nota < 5:
			return "D"
		case nota >= 0 && nota < 3:
			return "E"
		default:
			return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(11.8))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(2.8))

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func ()  {
		
	}))
	fmt.Println(tipo(time.Now()))
}