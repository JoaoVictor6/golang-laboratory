package main

import (
	"errors"
	"fmt"
)

var ErrorSaldoInsuficiente = errors.New("não é possivel retirar: saldo insuficiente")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

// quando criamos a carteira geramos uma cópia,
// para consertar isso utilizamos de ponteiros, basicamente um * na frete do "tipo"
//A diferença é que o tipo do argumento é *Carteira em vez de Carteira que você pode ler como "um ponteiro para uma carteira".

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade <= c.saldo {
		c.saldo -= quantidade
		return nil
	} else {
		return ErrorSaldoInsuficiente
	}
}

func main() {
}
