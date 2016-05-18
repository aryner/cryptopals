package main

import (
	"fmt"
)

var reducingPoly = int(283)

func peasantsMult(a int, b int) int {
	var p = int(0)
	for b > 0 {
		if b % 2 == 1 {
			p = p ^ a
		}
		b = b >> 1
		a = a << 1
		if a >= 256 {
			a = a ^ reducingPoly
		}
	}
	return p
}

func main() {
	fmt.Printf("%b\n",peasantsMult(83,202))
}
