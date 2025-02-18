package bloomfilter

import (
	"strconv"
	"testing"
)

func TestNeverFalseNegative(t *testing.T) {
	b := New(1_000_000)
	for i := 0; i < 1_000_000; i++ {
		b.Add("some" + strconv.Itoa(i))
	}
	for i := 0; i < 1_000_000; i++ {
		key := "some" + strconv.Itoa(i)
		if b.Has(key) != true {
			t.Errorf("b.Has('some%s') = %t, want true", strconv.Itoa(i), b.Has(key))
		}
	}
}

func TestFalsePositiveRate(t *testing.T) {
	n := 1_000_000
	b := New(n)
	for i := 0; i < n; i++ {
		b.Add("some" + strconv.Itoa(i))
	}

	falsePositives := 0
	testSize := 1_000_000

	for i := 0; i < testSize; i++ {
		key := "not_added" + strconv.Itoa(i)
		if b.Has(key) {
			falsePositives++
		}
	}

	rate := float64(falsePositives) / float64(testSize)
	if rate > 0.02 { // Allowing a bit of margin
		t.Errorf("False positive rate too high: got %f, want ≤ 0.01", rate)
	}
}
