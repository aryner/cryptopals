package main

import (
        "fmt"
	"encoding/hex"
)

func SingleByteXORCipher(c byte, b []byte) []byte {
	var result []byte
	for i := range b {
		result = append(result, c ^b[i])
	}
	return result
}

func SingleByteXORDecode(coded []byte) {
	for i:=0; i<128; i++ {
		decode := SingleByteXORCipher(byte(i),coded)
		fmt.Println(fmt.Sprintf("%c",byte(i)))
		for _, v := range decode {
			fmt.Printf("%c",v)
		}
		fmt.Println()
	}
}

func main() {
        coded, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	SingleByteXORDecode(coded)
}

