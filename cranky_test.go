package cranky

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestPow10(t *testing.T) {
	for i := uint64(0); i < 15; i++ {
		t.Logf("%d - %d", i, pow10(i))
	}
}

func TestPermute(t *testing.T) {
	for p := range permute([]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Logf("%v", p)
	}
}

func TestDigitize(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := rand.Uint64()
		d := digitize(v)
		u := undigitize(d)
		t.Logf("%d - %v - %d", v, d, u)

		if v != u {
			t.Fatalf("input (%d) does not equal output (%d)", v, u)
		}
	}
}

func TestSumCranky(t *testing.T) {
	if sum, expected := SumCranky(runtime.NumCPU(), 0, pow10(6)), uint64(1778723); sum != expected {
		t.Logf("sum is %d; expected %d", sum, expected)
	}
}

func TestBigCalculation(t *testing.T) {
	sum := SumCranky(runtime.NumCPU(), 0, pow10(14))
	t.Logf("answer = %d", sum)
}
