// The Homophonic cipher is a substitution cipher that replaces each character
// with a randomly selected character from a set of possible substitutions.
// This makes frequency analysis more challenging and adds an extra layer of security.
package ciphers

import (
	"math/rand"
	"time"
)

func (c *CiphersController) CodeHomophonic(inputString string) string {
	var output string
	rand.Seed(time.Now().UnixNano())

	for _, character := range inputString {
		index := rand.Intn(len(c.homophonicMap[character]))

		output += string(c.homophonicMap[character][index])
	}

	return output
}

func (c *CiphersController) DecodeHomophonic(inputString string) string {
	var output string

	reverseMap := make(map[rune]rune)

	for key, runes := range c.homophonicMap {
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
