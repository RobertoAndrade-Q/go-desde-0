package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1500.66
	Fecha = time.Now()

	fmt.Println(Nombre, Estado, Sueldo, Fecha)
}

func ConverToText(n int) (bool, string) {
	var texto string
	texto = strconv.Itoa(n)

	return true, texto
}
