package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("Coleção de 5 numeros", func(t *testing.T) {
		arrayMock := []int{1, 2, 3, 4, 5}
		result := Soma(arrayMock)
		expected := 15

		if result != expected {
			t.Errorf("Expected: %d\nResult: %d", expected, result)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	verificaSomas := func(t *testing.T, result, expected []int) {
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result: %d\nExpected: %d", result, expected)
		}
	}

	t.Run("Faz a soma de alguns slices", func(t *testing.T) {
		result := SomaTudo([]int{1, 2}, []int{3, 9})
		expected := []int{3, 12}

		verificaSomas(t, result, expected)
	})

	t.Run("Soma slices vazios de forma seguran", func(t *testing.T) {
		result := SomaTudo([]int{}, []int{3, 9})
		expected := []int{0, 12}

		verificaSomas(t, result, expected)
	})
}
