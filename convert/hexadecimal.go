package convert

import "fmt"

// Hexadecimal converts a decimal integer to Hexadecimal
func Hexadecimal(i int) {
	Print("Dec-Hex", fmt.Sprintf("%d", i), fmt.Sprintf("%x", i))
}
