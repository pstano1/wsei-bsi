package ciphers

import (
	"github.com/sirupsen/logrus"
)

type ICiphersController interface {
	CipherCaesar(inputString string, shift int) string
	DecipherCaesar(inputString string, shift int) string
}

type CiphersController struct {
	log          logrus.FieldLogger
	characterSet []rune
}

func NewCiphersController(log logrus.FieldLogger, characterSet []rune) ICiphersController {
	return &CiphersController{
		log:          log,
		characterSet: characterSet,
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
