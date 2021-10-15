package main

import "time"

// ----> TESTES STRUCTS AND INTERFACES <----

type tempoSpy struct {
	duracaoPausa time.Duration
}

func (t *tempoSpy) Sleep(duracao time.Duration) {
	t.duracaoPausa = duracao
}

const (
	pausa   = "pausa"
	escrita = "escrita"
)

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type SleeperSpy struct {
	chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.chamadas++
}

// ----> DEFAULT STRUCTS AND INTERFACES <----
type SleeperConfiguravel struct {
	duracao time.Duration
	sleep   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	time.Sleep(s.duracao * time.Second)
}

type Sleeper interface {
	Sleep()
}
