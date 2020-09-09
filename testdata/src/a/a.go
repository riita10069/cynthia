package a

import "testing"

type A struct {}

func (this *A)F(x int) int {
	return x * x
}

func (this *A)G(x int) int { // want "not implemented"
	return x * x
}

func TestF(t *testing.T) {
	a := A{}
	patterns := []struct {
		x        int
		expected int
	}{
		{1, 1},
		{10, 100},
		{-10, 100},
	}

	for _, pattern := range patterns {
		actual := a.F(pattern.x)
		if pattern.expected != actual {
			t.Errorf("error!!")
		}
	}
}
