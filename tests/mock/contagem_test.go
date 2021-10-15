package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestContagem(t *testing.T) {
	t.Run("Imprime até 3 e vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}

		Contagem(buffer, sleeperSpy)

		sleeperCallsExpected := 4
		sleeperCalls := sleeperSpy.chamadas

		resultado := buffer.String()
		esperado := "3\n2\n1\nVai!\n"

		if resultado != esperado {
			t.Errorf("\nResultado: %v\nEsperado: %v", resultado, esperado)
		}
		if sleeperSpy.chamadas != sleeperCallsExpected {
			t.Errorf("Não houve a quantidade esperada de chamadas no sleeper.\nEsperado: %v\nResultado: %v", sleeperCallsExpected, sleeperCalls)
		}
	})

	t.Run("Pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("\nEsperado: %v\nResultado: %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}
