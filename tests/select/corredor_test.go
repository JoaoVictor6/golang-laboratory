package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("compara a velocidade e retorna o mais rápido", func(t *testing.T) {
		servidorLento := serverTest(20 * time.Microsecond)
		servidorRapido := serverTest(0 * time.Nanosecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL
		esperado := URLRapida

		resultado, err := Corredor(URLLenta, URLRapida)

		if err != nil {
			t.Fatalf("Não esperavamos um erro, porém obtivemos.\nError: %v", err)
		}

		if resultado != esperado {
			t.Errorf("\nresultado: %v\nesperado: %v", resultado, esperado)
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidor := serverTest(10 * time.Second)

		defer servidor.Close()

		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("Esperava um erro, porém não obtivemos")
		}
	})
}

func serverTest(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
