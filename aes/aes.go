package main

import (
	"fmt"
)

var reducingPoly = int(283)

//This example has cache, timing, and branch prediction side-channel leaks, and is not suitable for use in cryptography.
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

//Rotates a 32 bit word 8 bits to the left, wrapping back to the begining
func Rotate(word int) int {
	tail := (word & 0xff000000) >> 24
	rotated := ((word << 8) & 0xffffff00) ^ tail
	return rotated
}

func Rcon(i int) int {
	rcon := 1
	i--
	for ; i > 0; i-- {
		rcon = rcon << 1
		if rcon >= 256 {
			rcon ^=  reducingPoly
		}
	}
	return rcon
}

func GetRconTo(i int) []int {
	rcon := make([]int, i, i)
	for ;i>0; i-- {
		rcon[i-1] = Rcon(i)
	}
	return rcon
}

func main() {
	fmt.Printf("%b\n",peasantsMult(83,202))
	fmt.Printf("%x\n",Rotate(0x1d2c3a4f))
	rcon := GetRconTo(10)
	for i, v := range rcon {
		fmt.Printf("%d - %x\n",i+1,v)
	}
}
