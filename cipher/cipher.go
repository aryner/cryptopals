package main

import (
        "fmt"
	"encoding/hex"
)

var LettersNumbers []byte

func initLetters() {
	if len(LettersNumbers) == 0 {
		for i:='A'; i<='Z'; i++ {
			LettersNumbers = append(LettersNumbers, byte(i))
		}
		for i:='a'; i<='z'; i++ {
			LettersNumbers = append(LettersNumbers, byte(i))
		}
		for i:=0; i<10; i++ {
			LettersNumbers = append(LettersNumbers, byte(i))
		}
	}
}

func SingleByteXORCipher(c byte, b []byte) []byte {
	var result []byte
	for i := range b {
		result = append(result, c ^b[i])
	}
	return result
}

func SingleByteXORDecode(coded []byte) {
	initLetters()
	m := make(map[byte][]byte)
	for _, c := range LettersNumbers {
		decode := SingleByteXORCipher(c,coded)
		m[c] = decode
		/*
		fmt.Println(fmt.Sprintf("%c",byte(i)))
		for _, v := range decode {
			fmt.Printf("%c",v)
		}
		fmt.Println()
		*/
	}
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%c",k))
		for _, c := range v {
			fmt.Printf("%c",c)
		}
		fmt.Println()
	}
}

func main() {
        coded, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	SingleByteXORDecode(coded)
}

