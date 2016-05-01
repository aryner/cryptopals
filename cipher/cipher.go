package main

import (
        "fmt"
	"encoding/hex"
)

var LettersNumbers []byte
var LetterFrequency = map[byte]float64 {
	'E':12.02, 'T':9.1, 'A':8.12, 'O':7.68, 'I':7.31,
	'N':6.95, 'S':6.28, 'R':6.02, 'H':5.92, 'D':4.32,
	'L':3.98, 'U':2.88, 'C':2.71, 'M':2.61, 'F':2.3,
	'Y':2.11, 'W':2.09, 'G':2.03, 'P':1.82, 'B':1.49,
	'V':1.11, 'K':0.69, 'X':0.17, 'Q':0.11, 'J':0.1,
	'Z':0.07,
}

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
	}
	printDecodedMaps(m)
}

func printDecodedMaps(m map[byte][]byte) {
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

