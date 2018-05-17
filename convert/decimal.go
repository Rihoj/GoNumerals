package convert

import "fmt"

// Decimal converts a decimal integer to decimal
func Decimal(i int) {
	Print("Dec-Dec", fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
}
