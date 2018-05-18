package convert

import "fmt"

// Print prints the value on a new line in a standard format
func Print(label string, orig string, new string) {
	fmt.Printf("%8s\t%s\t%s\n", label, orig, new)
}

// All prints out all numeral conversions orders by base
func All(i int) {
	Print("Label", "Orig", "Converted")
	Binary(i)
	Octal(i)
	Decimal(i)
	Hexadecimal(i)
	Utf8(i)
	fmt.Println()
}
