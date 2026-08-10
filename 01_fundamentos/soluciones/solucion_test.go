package solucion

import "testing"

func TestCalcularTotal(t *testing.T) {
	total, err := CalcularTotal(250, 3)
	if err != nil || total != 750 {
		t.Fatalf("total=%d err=%v", total, err)
	}
}
