package main

import (
	"fmt"

	"github.com/Rihoj/goNumerals/convert"
)

func main() {
	for i := 0; i < 17; i++ {
		convert.Decimal(i)
		convert.Binary(i)
		convert.Octal(i)
		convert.Hexadecimal(i)
		fmt.Println(" ")
	}
}
