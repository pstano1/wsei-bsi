package ciphers

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

type ICiphersController interface {
	CodeCaesar(inputString string, shift int) string
	DecodeCaesar(inputString string, shift int) string
	CodePolybiusSquare(inputString string) string
	DecodePolybiusSquare(inputString string) string
	CodeBeale(inputString string) string
	DecodeBeale(inputString string) string
	CodeTrithemius(inputString string, key rune) string
	DecodeTrithemius(inputString string, key rune) string

	ClearInput(input string) string
	searchForRune(character rune, characterSet []rune) int
	searchForRune2D(character rune, characterSet [][]rune) (int, int)
	polybiusMapping(index int, operation string) int
}

type CiphersController struct {
	log            logrus.FieldLogger
	characterSet   []rune
	polybiusSquare [][]rune
	bealeMap       map[rune][]rune
}

func NewCiphersController(log logrus.FieldLogger, characterSet []rune, polybiusSquare [][]rune, bealeMap map[rune][]rune) ICiphersController {
	return &CiphersController{
		log:            log,
		characterSet:   characterSet,
		polybiusSquare: polybiusSquare,
		bealeMap:       bealeMap,
	}
}

func (c *CiphersController) searchForRune(character rune, characterSet []rune) int {
	index := -1
	for i, r := range characterSet {
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

func (c *CiphersController) ClearInput(input string) string {
	input = strings.ReplaceAll(input, " ", "")

	digitRe := regexp.MustCompile("[0-9]")
	input = digitRe.ReplaceAllString(input, "")

	punctRe := regexp.MustCompile(`[[:punct:]]`)
	input = punctRe.ReplaceAllString(input, "")

	return strings.ToLower(input)
}
