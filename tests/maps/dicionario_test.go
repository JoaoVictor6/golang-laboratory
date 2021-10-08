package main

import (
	"testing"
)

func comparaErro(t *testing.T, erro error, messageError error) {
	t.Helper()

	if erro != messageError {
		t.Errorf("\nError: %s\nError esperado: %s", erro, messageError)
	}
}
func TestBusca(t *testing.T) {
	comparaStrings := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("\nresultado: %s\nesperado: %s", resultado, esperado)
		}
	}

	t.Run("Palavra existente", func(t *testing.T) {
		const esperado string = "isso é apenas um teste"
		dicionario := Dicionario{"teste": "isso é apenas um teste"}

		resultado, _ := dicionario.Busca("teste")

		comparaStrings(t, resultado, esperado)
	})

	t.Run("Palavra inexistente", func(t *testing.T) {
		const palavraChave string = "INEXISTENTE"
		dicionario := Dicionario{"teste": "isso é apenas um teste"}
		_, err := dicionario.Busca(palavraChave)

		comparaErro(t, err, ErrNaoEncontrado)
	})
}

func comparaDefinicoes(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()
	resultado, err := dicionario.Busca(palavra)

	if err != nil {
		t.Fatal("Não foi possivel achar a palavra adicionada")
	}
	if resultado != definicao {
		t.Errorf("\nResultado: %s\nEsperado: %s", resultado, definicao)
	}
}
func TestAdiciona(t *testing.T) {

	t.Run("Adicionando nova palavra", func(t *testing.T) {
		dicionario := Dicionario{}
		const definicao string = "Apenas um teste"
		const palavra string = "teste"
		err := dicionario.Adiciona(palavra, definicao)

		comparaErro(t, err, nil)
		comparaDefinicoes(t, dicionario, palavra, definicao)
	})

	t.Run("Adiciona palavra já existente", func(t *testing.T) {
		const palavra string = "teste"
		const definicao string = "Apenas um teste"
		dicionario := Dicionario{palavra: definicao}

		// Adicionando de novo
		err := dicionario.Adiciona(palavra, "e")

		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicoes(t, dicionario, palavra, definicao)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Atualizando valor de uma chave existente", func(t *testing.T) {
		const palavra string = "teste"
		const definicao = "Apenas um teste"
		const novaDefinicao string = "Novadefinição"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualizar(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicoes(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("Atualizando valor de uma chave inexistente", func(t *testing.T) {
		const palavra string = "teste"
		const definicao = "Apenas um teste"

		const novaDefinicao string = "Novadefinição"
		const palavraInexistente string = "aaaaa"

		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualizar(palavraInexistente, novaDefinicao)

		comparaErro(t, err, ErrPalavraInexistente)
	})
}
