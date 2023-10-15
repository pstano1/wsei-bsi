package ciphers

import (
	"github.com/sirupsen/logrus"
)

type ICiphersController interface {
	CodeCaesar(inputString string, shift int) string
	DecodeCaesar(inputString string, shift int) string
	CodePolybiusSquare(inputString string) string
	DecodePolybiusSquare(inputString string) string
}

type CiphersController struct {
	log            logrus.FieldLogger
	characterSet   []rune
	polybiusSquare [][]rune
}

func NewCiphersController(log logrus.FieldLogger, characterSet []rune, polybiusSquare [][]rune) ICiphersController {
	return &CiphersController{
		log:            log,
		characterSet:   characterSet,
		polybiusSquare: polybiusSquare,
	}
}

func (c *CiphersController) searchForRune(character rune) int {
	index := -1
	for i, r := range c.characterSet {
		if r == character {
			index = i
			break
		}
	}

	return index
}

func (c *CiphersController) searchForRune2D(character rune, characterSet [][]rune) (int, int) {
	for row, rowCharacters := range characterSet {
		for index, r := range rowCharacters {
			if r == character {
				return row, index
			}
		}
	}

	return -1, -1
}
