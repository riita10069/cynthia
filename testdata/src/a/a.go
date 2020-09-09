package a

type A struct {}

func (this *A)F(x int) int {
	return x * x
}

func (this *A)G(x int) int { // want "not implemented"
	return x * x
}
