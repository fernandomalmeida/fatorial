package fatorial

import (
	"errors"
)

func Fatorial(n uint64) (uint64, error) {
	if n > 20 {
		return 0, errors.New("Não pode calcular o fatorial de um número maior que 20")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}

	var resultado uint64
	resultado = 1
	for i := n; i > 1; i-- {
		resultado *= i
	}

	return resultado, nil
}
