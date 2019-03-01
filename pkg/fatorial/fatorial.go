package fatorial

import (
	"errors"
    "fmt"
)

const limiteSuperior = 20

func Fatorial(n uint64) (uint64, error) {
	if n > limiteSuperior {
		return 0, errors.New(fmt.Sprintf("Não pode calcular o fatorial de um número maior que %d", limiteSuperior))
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
