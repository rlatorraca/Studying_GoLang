package math

var A int = 10
var a int = 15

type Car struct {
	Make  string
	Model string
}

func (c Car) Run() string {
	return "Car running"
}

// Sum
/* o nome da funcao precisar estar em MAISCULO para ser vista de FORA do pacote*/
func Sum[T int | float64](x, y T) T {
	return x + y
}
