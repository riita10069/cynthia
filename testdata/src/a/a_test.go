package a

import "testing"

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
