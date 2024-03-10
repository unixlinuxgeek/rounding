package rounding

import (
	"math/rand"
	"testing"
)

func TestRound(t *testing.T) {
	r := 1 + rand.Float64()*(9-1) // generate random int (1 - 9)
	t.Logf("Generate random float64 number(1-9): %f\n", r)

	if Round(r) == r {
		t.Errorf("Test Failed: %f  is equal %f", r, Round(r))
	} else {
		t.Log("Test Round Passed!!!")
	}
}
