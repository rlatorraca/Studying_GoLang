package main

const letra_a = "Hello World A"

var letra_b bool

var (
	var01 int
	var02 string
	var03 float32
	var04 float64
	var05 byte
)

var (
	var06 int     = 123
	var07 string  = "Testing"
	var08 float32 = 1.33
	var09 float64 = 2.344
	var10 byte
)

func main() {
	println(letra_a)
	println(letra_b)
	println(var01)
	println(var02)
	println(var03)
	println(var04)
	println(var06)
	println(var07)
	println(var08)
	println(var09)

	new_var_por_inferencia := "Sem tipo" // usando := na primeira vez
	println(new_var_por_inferencia)
	new_var_por_inferencia = "Sem tipo, modificado" // nao usa : da segunda vez em diante
	println(new_var_por_inferencia)
}
