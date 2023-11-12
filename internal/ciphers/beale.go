package ciphers

import (
	"math/rand"
	"time"
)

func (c *CiphersController) CodeBeale(inputString string) string {
	var output string

	for _, character := range inputString {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(c.bealeMap[character]))

		output += string(c.bealeMap[character][index])
	}

	return output
}

func (c *CiphersController) DecodeBeale(inputString string) string {
	var output string

	reverseMap := make(map[rune]rune)

	for key, runes := range c.bealeMap {
		for _, r := range runes {
			reverseMap[r] = key
		}
	}

	for _, character := range inputString {
		decodedRune, exists := reverseMap[character]
		if exists {
			output += string(decodedRune)
		} else {
			output += string(character)
		}
	}

	return output
}
