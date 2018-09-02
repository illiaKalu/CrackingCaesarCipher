package utils

import (
	"strings"
	"unicode"
)

// English alphabet
var alphabetArray = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}


func CalculateGivenTextFrequency(text string) map[string]float32 {

	freqMap := make(map[string]float32)

	// remove all spaces and stuff
	text = removeBadCasesFromText(text)

	for i := 0; i < len(alphabetArray); i++ {
		freqMap[alphabetArray[i]] = float32(strings.Count(text, alphabetArray[i])) * 100.0 / float32(len(text))
	}
	return freqMap
}

// Removes all chars that doesn't take action in frequency calculation from incoming text
func removeBadCasesFromText(text string) (string) {

	text = strings.ToLower(text)
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, "'", "", -1)
	text = strings.Replace(text, "\t", "", -1)

	return text
}

/* Calculates distance between
/ Main algorithm
/ Early build, for now gets top frequent letter occurrence from text and
/ measure distance between it and 'e' ( most frequent letter in Eng alphabet )
*/
func GetMovedFactor(currentLetterFrequency map[string]float32, letterFrequency map[string]float32) (z int) {
	var max float32 = -1
	var letter string
	for key, val := range currentLetterFrequency {
		if max < val {
			max = val
			letter = key
		}
	}
	z = int('e' - ([]rune(letter)[0]))
	return
}

// Shuffles text by Z number of positions
func ShuffleText(text string, z int) string {
	textR := []rune(text)
	for i := 0; i < len(text); i++ {
		if unicode.IsLetter(textR[i]) {
			textR[i] += rune(z)
		}
	}

	return string(textR)
}

func FillMap(letterFrequency map[string]float32) {
	// English alphabet letter frequency in %
	letterFrequency["a"] = 8.167
	letterFrequency["b"] = 1.492
	letterFrequency["c"] = 2.782
	letterFrequency["d"] = 4.253
	letterFrequency["e"] = 12.70
	letterFrequency["f"] = 2.228
	letterFrequency["g"] = 2.015
	letterFrequency["h"] = 6.094
	letterFrequency["i"] = 6.966
	letterFrequency["j"] = 0.153
	letterFrequency["k"] = 0.772
	letterFrequency["l"] = 4.025
	letterFrequency["m"] = 2.406
	letterFrequency["n"] = 6.749
	letterFrequency["o"] = 7.507
	letterFrequency["p"] = 1.929
	letterFrequency["q"] = 0.095
	letterFrequency["r"] = 5.987
	letterFrequency["s"] = 6.327
	letterFrequency["t"] = 9.056
	letterFrequency["u"] = 2.758
	letterFrequency["v"] = 0.978
	letterFrequency["w"] = 2.360
	letterFrequency["x"] = 0.150
	letterFrequency["y"] = 1.974
	letterFrequency["z"] = 0.074
}
