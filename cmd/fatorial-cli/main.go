package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fernandomalmeida/fatorial/pkg/fatorial"
)

const (
	errUtilizacao = iota + 1
	errParse
	errFatorial
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("A utilização do programa deve ser:\n%s n\tSendo n o valor a ter o fatorial calculado.\n", os.Args[0])
		os.Exit(errUtilizacao)
	}
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Erro ao obter o valor de n: %s\n", err)
		os.Exit(errParse)
	}
	resultado, err := fatorial.Fatorial(n)
	if err != nil {
		fmt.Printf("Erro ao calcular o fatorial: %s\n", err)
		os.Exit(errFatorial)
	}

	fmt.Println(resultado)
}
