package main

import (
	//"github.com/aryner/cryptopals/set1"
	"fmt"
)

func pad(b []byte, bytesize int) []byte {
	var padding []byte
	if mod := len(b) % bytesize; mod != 0 {
		for i:=0; i<bytesize-mod; i++ {
			padding = append(padding, 4)
		}
	}
	return append(b,padding...)
}

func main() {
	base := []byte("YELLOW SUBMARINE")
	padded := pad(base,20)
	testcase := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	fmt.Println(string(padded))
	fmt.Println(string(padded) == string(testcase))
}
