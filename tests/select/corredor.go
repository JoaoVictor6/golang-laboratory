package main

import (
	"fmt"
	"net/http"
	"time"
)

const limiteDezSegundos = 10 * time.Second

func Corredor(a, b string) (vencedor string, erro error) {
	return Configuravel(a, b, limiteDezSegundos)
}

func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, erro error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(tempoLimite):
		return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func(u string) {
		http.Get(u)
		ch <- true
	}(URL)

	return ch
}
