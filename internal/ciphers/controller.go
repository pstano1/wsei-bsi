package ciphers

import (
	"errors"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

type ICiphersController interface {
	CodeCaesar(inputString string, shift int) string
	DecodeCaesar(inputString string, shift int) string
	CodePolybiusSquare(inputString string) string
	DecodePolybiusSquare(inputString string) string
	CodeHomophonic(inputString string) string
	DecodeHomophonic(inputString string) string
	CodeTrithemius(inputString string, key rune) string
	DecodeTrithemius(inputString string, key rune) string
	CodeVigenere(inputString string, key string) string
	DecodeVigenere(inputString string, key string) string

	ClearInput(input string) string
	searchForRune(character rune, characterSet []rune) int
	searchForRune2D(character rune, characterSet [][]rune) (int, int)
	polybiusMapping(index int, operation string) int
}

type CiphersController struct {
	log            logrus.FieldLogger
	characterSet   []rune
	polybiusSquare [][]rune
	homophonicMap  map[rune][]rune
}

func NewCiphersController(log logrus.FieldLogger, characterSet []rune, polybiusSquare [][]rune, homophonicMap map[rune][]rune) (ICiphersController, error) {
	if hasDuplicate := checkHomophonicMapForDuplicates(homophonicMap); !hasDuplicate {
		return nil, errors.New("rune map for homophonic cipher has duplicates")
	}

	return &CiphersController{
		log:            log,
		characterSet:   characterSet,
		polybiusSquare: polybiusSquare,
		homophonicMap:  homophonicMap,
	}, nil
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

func checkHomophonicMapForDuplicates(homophonicMap map[rune][]rune) bool {
	seenValues := make(map[rune]struct{})

	for _, values := range homophonicMap {
		for _, value := range values {
			if _, exists := seenValues[value]; exists {
				return true
			}
			seenValues[value] = struct{}{}
		}
	}

	return false
}
