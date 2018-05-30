package fatorial

import "testing"

func TestFatorial(t *testing.T) {
	entradas := []struct {
		N         uint64
		Resultado uint64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, caso := range entradas {
		r, _ := Fatorial(caso.N)
		if r != caso.Resultado {
			t.Errorf("%d! deveria ser %d, mas resultou em %d", caso.N, caso.Resultado, r)
		}
	}
}
