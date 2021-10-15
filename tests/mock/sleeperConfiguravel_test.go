package main

import (
	"testing"
	"time"
)

func TestSleeperConfigurabel(t *testing.T) {
	tempoPausa := 5 * time.Second

	tempoSpy := &tempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Sleep}
	sleeper.Sleep()
	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}
}
