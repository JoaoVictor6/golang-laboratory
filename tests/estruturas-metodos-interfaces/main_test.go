package main

import (
	"testing"
)

//Suponha que precisamos de algum código
//de geometria para calcular o perímetro
//de um retângulo dado uma altura e largura.
//Podemos escrever uma função
// func verifyValues(t *testing.T, forma Forma, expected float64) {
// 	t.Helper()
// 	result := forma.Area()
// 	if result != expected {
// 		t.Errorf("resultado %.2f esperado %.2f", result, expected)
// 	}
// }
func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	const expected float64 = 40.0
	result := Perimetro(retangulo)

	if result != expected {
		t.Errorf("resultado %.2f esperado %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	const expectedCircleArea float64 = 314.1592653589793
	//Table drive test serve para quando se quer testar coisas que seguem um padrão.
	// Nesse caso o padrão é o mesmo, chamar o mesmo método
	testeArea := []struct {
		forma   Forma
		temArea float64
	}{
		{forma: Retangulo{Altura: 10.0, Largura: 10.0}, temArea: 100.0},
		{forma: Circulo{Raio: 10}, temArea: expectedCircleArea},
		{forma: Triangulo{Base: 12, Altura: 6}, temArea: 36},
	}

	for _, tt := range testeArea {
		result := tt.forma.Area()
		if result != tt.temArea {
			t.Errorf("\n%#v\nresultado %.2f\nesperado %.2f", tt.forma, result, tt.temArea)
		}
	}

}
