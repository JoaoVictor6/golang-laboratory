package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &SleeperConfiguravel{duracao: 3, sleep: time.Sleep}
	Contagem(os.Stdout, sleeper)
}

func Contagem(saida io.Writer, timeInstance Sleeper) {
	const inicioContagem = 3
	const ultimaPalavra = "Vai!"
	for i := inicioContagem; i > 0; i-- {
		timeInstance.Sleep()
		fmt.Fprintln(saida, i)
	}
	timeInstance.Sleep()
	fmt.Fprintln(saida, ultimaPalavra)
}
