package main

import (
	"fmt"
)

import (
	util "CrackingCaesarCipher/CrackingCaesarCipher/utils"
	"io/ioutil"
)

func main() {

	var letterFrequency = make(map[string]float32)
	var currentTextLetterFrequency = make(map[string]float32)
	// ROT1 - rotated text by one char by Caesar's cipher
	cipheredText, err := ioutil.ReadFile("ROT1.txt")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("text - ", string(cipheredText))
	}

	// fills the default english letters frequency map
	util.FillMap(letterFrequency)

	// calculates incoming text letter frequency and Z index (number of rotates)
	currentTextLetterFrequency = util.CalculateGivenTextFrequency(string(cipheredText))
	z := util.GetMovedFactor(currentTextLetterFrequency, letterFrequency)

	unCipheredText := util.ShuffleText(string(cipheredText), z)

	//fmt.Println("letters freq: ", letterFrequency)
	//fmt.Println("current text freq: ", currentTextLetterFrequency)

	fmt.Println("result: ", unCipheredText)

}


