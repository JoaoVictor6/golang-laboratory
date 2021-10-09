package main

import (
	"fmt"
)

const (
	ErrNaoEncontrado      = ErrDicionario("não foi possível encontrar a palavra que você procura")
	ErrPalavraExistente   = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErrDicionario("essa chave não existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

type Dicionario map[string]string

func (d Dicionario) Busca(palavraChave string) (string, error) {
	palavra, existe := d[palavraChave]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return palavra, nil
}

func (d Dicionario) Delete(palavraChave string) error {
	_, err := d.Busca(palavraChave)
	if err != nil {
		return ErrPalavraInexistente
	}
	delete(d, palavraChave)
	return nil
}

func (d Dicionario) Adiciona(palavraChave, palavra string) error {
	_, existe := d[palavraChave]
	if existe {
		return ErrPalavraExistente
	}
	d[palavraChave] = palavra
	return nil
}

func (d Dicionario) Atualizar(palavraChave, definicao string) error {
	_, existe := d[palavraChave]
	if !existe {
		return ErrPalavraInexistente
	}
	d[palavraChave] = definicao
	return nil
}

func main() {
	fmt.Println("vim-go")
}
