package main

import (
        "fmt"
	"encoding/hex"
	"strings"
	"io/ioutil"
)

var Thresh = 5.5
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
	byte(39):0.0,
	byte(44):0.0,
}
var InvalidChar = -10.00
var CommonWords = map[string]float64 {
	"the":7.14, "of":4.16, "and":30.4, "to":2.6,
	"in":2.27, "a":2.06, "is":1.13, "that":1.08,
	"for":0.88, "it":0.77, "as":0.77, "was":0.74,
	"with":0.77, "be":0.65, "by":0.63, "on":0.62,
	"not":0.61, "he":0.55, "i":0.52, "this":0.51,
	"are":0.50, "or":0.49, "his":0.49, "from":0.47,
	"at":0.46, "which":0.42, "but":0.38, "have":0.37,
	"an":0.37, "had":0.35, "they":0.33, "you":0.31,
	"were":0.31, "their":0.29, "one":0.29, "all":0.28,
	"we":0.28, "can":0.22, "her":0.22, "has":0.22,
	"there":0.22, "been":0.22, "if":0.21, "more":0.21,
	"when":0.2, "will":0.2, "would":0.2, "who":0.2,
	"so":0.19, "no":0.19,"party":10.0,
}
var WordLengthFrequency = map[int]float64 {
	1:2.998, 2:17.651, 3:20.511, 4:14.787,
	5:10.7, 6:8.388, 7:7.939, 8:5.943,
	9:4.437, 10:3.076, 11:1.761, 12:0.958,
	13:0.518, 14:0.22, 15:0.076, 16:0.02,
	17:0.01, 18:0.004, 19:0.001, 20:0.001,
}
var TooLong = -0.50

func initLetters() {
	if len(LettersNumbers) == 0 {
		for i:='A'; i<='Z'; i++ {
			LettersNumbers = append(LettersNumbers, byte(i))
		}
		for i:='a'; i<='z'; i++ {
			LettersNumbers = append(LettersNumbers, byte(i))
		}
		for i:='0'; i<='9'; i++ {
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

func RepeatingKeyXORCipher(k []byte, m string) string {
	var result []byte
	for i := range m {
		result = append(result, k[i % len(k)] ^ byte(m[i]))
	}
	return hex.EncodeToString(result)
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

func ProposeKeyLength(min int, max int, text []byte) []int {
	var lengths []int
	var distances []float64
	for ; min<=max; min++ {
		dis := getAvgDistance(min,4,text)
		distances = append(distances, dis)
		lengths = append(lengths,min)
		distances, lengths = orderPairs(3,distances,lengths)
	}
	return lengths
}

func getAvgDistance(keyLength int, numBlocks int, text []byte) float64 {
	var sum float64
	var n int
	for i:=0; i<numBlocks; i++ {
		for j:=i+1; j<=numBlocks; j++ {
			if keyLength*(j+1) < len(text) {
				sum += float64(HammingDistance(text[keyLength*i:keyLength*(i+1)],text[keyLength*j:keyLength*(j+1)])/keyLength)
				n++
			}
		}
	}
	return sum / float64(n)
}

func orderPairs(max int, sortBy []float64, follow []int) ([]float64, []int) {
	if len(sortBy) <= 1 {
		return sortBy, follow
	} else {
		back := len(sortBy)-1
		for ;back > 0 && sortBy[back] < sortBy[back-1]; back-- {
			temp := sortBy[back]
			sortBy[back] = sortBy[back-1]
			sortBy[back-1] = temp
			temp2 := follow[back]
			follow[back] = follow[back-1]
			follow[back-1] = temp2
		}
	}
	if len(sortBy) < max {
		return sortBy, follow
	}
	
	return sortBy[:max], follow[:max]
}

func HammingDistance(one []byte, two []byte) int {
	var result int
	for i, v := range one {
		diff := v ^ two[i]
		for ; diff > 0; diff = diff >> 1 {
			if diff % 2 != 0 {
				result += 1
			}
		}
	}
	return result
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
			result += InvalidChar
		}
	}
	result /= float64(len(b))
	result += ScoreCommonWords(b)
/*
	if result > 2 {
		result += ScoreOnWordLength(b)
	}
*/
	return result 
}

func ScoreCommonWords(b []byte) float64{
	var result float64
	words := strings.Split(string(b), " ")
	for _, word := range words {
		if v, ok := CommonWords[strings.ToLower(word)]; ok {
			result += v
		}
	}
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
	return result / float64(len(words))
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

func PossibleSingleByteXOR(b []byte) (bool,float64,byte,[]byte) {
	cipher, score, byteArray := SingleByteXORDecode(b)
	if score > Thresh {
		return true, score, cipher, byteArray
	}
	return false, score, cipher, byteArray
}

func DetectSingleByteXORs(codes []string) {
	for _, code := range codes {
		hexCode, _ := hex.DecodeString(code)
		passes,score,cipher,array := PossibleSingleByteXOR(hexCode)
		if passes {
			fmt.Println(fmt.Sprintf("%c - %v",cipher,score))
			printFormatedByteArray(array)
		}
	}
}

func main() {
//Test repeating xor decoding
	text, err := ioutil.ReadFile("test2.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		possibleLengths := ProposeKeyLength(2,40,text)
		fmt.Println(possibleLengths)
	}

//Test the Hamming distance function
	fmt.Println(HammingDistance([]byte("this is a test"),[]byte("wokka wokka!!!")))

//Test code for repeating XOR cipher
	var message = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := []byte{'I','C','E'}
	rxorCoded := RepeatingKeyXORCipher(key,message)
	fmt.Println(rxorCoded)

//Test code for detecting an single byte XOR cipher
	f, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		codes := strings.Split(string(f),"\n")
		DetectSingleByteXORs(codes)
	}

//Test code for single byte XOR cipher decoding
        coded, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	cipher, score, byteArray := SingleByteXORDecode(coded)
	fmt.Println(fmt.Sprintf("%c - %v",cipher,score))
	printFormatedByteArray(byteArray)
}

