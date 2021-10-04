package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola()
	expected := "Olá, mundo"

	if result != expected {
		t.Errorf("Resultado: %s\nExpected: %s", result, expected)
	}
}
