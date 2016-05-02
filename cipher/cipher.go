package main

import (
        "fmt"
	"encoding/hex"
	"strings"
)

var LettersNumbers []byte
var LetterFrequency = map[byte]float64 {
	'E':12.02, 'T':9.1, 'A':8.12, 'O':7.68, 'I':7.31,
	'N':6.95, 'S':6.28, 'R':6.02, 'H':5.92, 'D':4.32,
	'L':3.98, 'U':2.88, 'C':2.71, 'M':2.61, 'F':2.3,
	'Y':2.11, 'W':2.09, 'G':2.03, 'P':1.82, 'B':1.49,
	'V':1.11, 'K':0.69, 'X':0.17, 'Q':0.11, 'J':0.1,
	'Z':0.07,
	'e':12.02, 't':9.1, 'a':8.12, 'o':7.68, 'i':7.31,
	'n':6.95, 's':6.28, 'r':6.02, 'h':5.92, 'd':4.32,
	'l':3.98, 'u':2.88, 'c':2.71, 'm':2.61, 'f':2.3,
	'y':2.11, 'w':2.09, 'g':2.03, 'p':1.82, 'b':1.49,
	'v':1.11, 'k':0.69, 'x':0.17, 'q':0.11, 'j':0.1,
	'z':0.07,
	byte(32):0.0,
	byte(34):0.0,
	byte(44):0.0,
}
var InvalidChar = -0.15
var WordLengthFrequency = map[int]float64 {
	1:2.998, 2:17.651, 3:20.511, 4:14.787,
	5:10.7, 6:8.388, 7:7.939, 8:5.943,
	9:4.437, 10:3.076, 11:1.761, 12:0.958,
	13:0.518, 14:0.22, 15:0.076, 16:0.02,
	17:0.01, 18:0.004, 19:0.001, 20:0.001,
}
var TooLong = -0.01

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

func SingleByteXORDecode(coded []byte) (byte, float64, []byte) {
	initLetters()
	m := make(map[byte][]byte)
	for _, c := range LettersNumbers {
		decode := SingleByteXORCipher(c,coded)
		m[c] = decode
	}

	scores := ScoreMaps(m)
	k := GetHighScoreCipher(scores,m)
	return k, scores[k], m[k]
}

func ScoreMaps(m map[byte][]byte) map[byte]float64 {
	var result = make(map[byte]float64)
	for k, v := range m {
		result[k] = ScoreAsEnglish(v)
	}
	return result
}

func ScoreAsEnglish(b []byte) float64{
	var result float64
	for _, c := range b {
		if v, ok := LetterFrequency[c]; ok {
			result += v
		} else {
			result -= InvalidChar
		}
	}
	result /= float64(len(b))
	result += ScoreOnWordLength(b)
	return result 
}

func ScoreOnWordLength(b []byte) float64{
	var result float64
	words := strings.Split(string(b), " ")
	for _, word := range words {
		if v, ok := WordLengthFrequency[len(word)]; ok {
			result += v
		} else {
			result += TooLong * float64((len(word) - 20))
		}
	}
	return result
}

func printHighScores(s map[byte]float64, m map[byte][]byte) {
	for k, v := range s {
		if v > 3.5 {
			fmt.Println(fmt.Sprintf("%c - %v",k,v))
			printFormatedByteArray(m[k])
		}
	}
}

func GetHighScoreCipher(s map[byte]float64, m map[byte][]byte) byte {
	highScore := 0.0
	var highCipher byte
	for k,v := range s {
		if v > highScore {
			highScore = v
			highCipher = k
		}
	}
	return highCipher
}

func printDecodedMaps(m map[byte][]byte) {
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%c",k))
		printFormatedByteArray(v)
	}
}

func printFormatedByteArray(v []byte) {
	for _, c := range v {
		fmt.Printf("%c",c)
	}
	fmt.Println()
}

func main() {
        coded, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	cipher, score, byteArray := SingleByteXORDecode(coded)
	fmt.Println(fmt.Sprintf("%c - %v",cipher,score))
	printFormatedByteArray(byteArray)
}

