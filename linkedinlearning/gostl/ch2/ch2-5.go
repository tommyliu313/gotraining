package main

/* import (
	"fmt"
	"strings"
)
*/

func main() {

	shift := 2
	s := "The quick brown fox jumps over the lazy dog"

	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') = (int(r) - int('A') + shift)
			if value > 91 {
				value -= 26
			} else if value < 65 {
				value += 26
			}
		}
	}
}
