package main

import (
	"fmt"

	"github.com/RobertoAndrade-Q/go-desde-0/variables"
)

func main() {
	// variables.ShowInts()
	// variables.RestVariables()
	estado, texto := variables.ConverToText(1000)

	fmt.Println(estado)
	fmt.Println(texto)

}
