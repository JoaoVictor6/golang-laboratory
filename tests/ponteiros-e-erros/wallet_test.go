package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	confirmaErroInexistente := func(t *testing.T, erro error) {
		if erro != nil {
			t.Errorf("\nErro inesperador recebido: \n\t%s", erro)
		}
	}

	confirmaErro := func(t *testing.T, erro error, expectedErrorMessage error) {
		t.Helper()

		if erro == nil {
			t.Error("Espera um erro mas nenhum ocorreu")
		}

		if erro != expectedErrorMessage {
			t.Errorf("\nResultado: %s\nEsperado: %s", erro, expectedErrorMessage)
		}
	}
	t.Run("Depositar", func(t *testing.T) {
		wallet := Carteira{saldo: 0}
		const esperado Bitcoin = 10

		wallet.Depositar(Bitcoin(10))
		confirmaSaldo(t, wallet, esperado)

	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		const esperado Bitcoin = 0
		wallet := Carteira{Bitcoin(10)}
		erro := wallet.Retirar(Bitcoin(10))

		confirmaSaldo(t, wallet, esperado)
		confirmaErroInexistente(t, erro)
	})

	t.Run("Retirar mais que o saldo", func(t *testing.T) {
		const saldoInicial Bitcoin = 0
		wallet := Carteira{saldoInicial}
		erro := wallet.Retirar(Bitcoin(100))

		confirmaSaldo(t, wallet, saldoInicial)
		confirmaErro(t, erro, ErrorSaldoInsuficiente)

	})
}
