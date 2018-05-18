package convert

import "fmt"

// Utf8 converts a decimal integer to octal
func Utf8(i int) {
	Print("Dec-UTF8", fmt.Sprintf("%d", i), fmt.Sprintf("%x", i))
}
