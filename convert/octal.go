package convert

import "fmt"

// Octal converts a decimal integer to octal
func Octal(i int) {
	Print("Dec-Oct", fmt.Sprintf("%d", i), fmt.Sprintf("%o", i))
}
