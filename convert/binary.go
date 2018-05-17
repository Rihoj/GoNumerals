package convert

import "fmt"

// Binary converts an decimal integer into binary
func Binary(i int) {
	Print("Dec-Bin", fmt.Sprintf("%d", i), fmt.Sprintf("%b", i))
}
