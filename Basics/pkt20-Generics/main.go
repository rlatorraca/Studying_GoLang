package main

type MyNumber uint8

// Number /* CONSTRAINTS */
type Number interface {
	~uint8 | uint16 | uint32 | float32 | float64
}

// Soma /* GENERICS */
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, value := range m {
		soma += value
	}
	return soma
}

// Compare /* Usa o pacote de CONSTRAINTS comparable */
func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m0 := map[string]uint16{"John": 7, "Michael": 22, "Jason": 44}
	m1 := map[string]uint16{"John": 741, "Michael": 852, "Jason": 963}
	m2 := map[string]uint16{"John": 7, "Michael": 22, "Jason": 44}
	m3 := map[string]float64{"John": 741.55, "Michael": 852.44, "Jason": 963.33}
	m4 := map[string]MyNumber{"John": 12, "Michael": 23, "Jason": 44}
	println(Soma(m1))
	println(Soma(m2))
	println(Soma(m3))
	println(Soma(m4))
	println(Compare(Soma(m0), Soma(m1)))
	println(Compare(Soma(m0), Soma(m2)))

}
