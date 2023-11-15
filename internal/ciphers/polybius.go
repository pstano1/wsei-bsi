// The Polybius Square cipher is a substitution cipher that represents each letter
// with a pair of coordinates in a grid. The encoding process involves mapping each
// character to its corresponding coordinates and converting them to a numeric value.
// This implementation also performs some math calcutaions to make it harder for analysis.
package ciphers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *CiphersController) CodePolybiusSquare(inputString string) string {
	var output string

	for _, character := range inputString {
		x, y := c.searchForRune2D(character, c.polybiusSquare)
		x = c.polybiusMapping(x, "code")
		y = c.polybiusMapping(y, "code")
		concatInt, _ := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))

		output += fmt.Sprintf("%d ", concatInt*concatInt)
	}

	return output
}

func (c *CiphersController) DecodePolybiusSquare(inputString string) string {
	var output string
	elements := strings.Split(inputString, " ")
	if len(strings.TrimSpace(inputString)) == 0 {
		return ""
	}

	for _, substr := range elements {
		number, err := strconv.Atoi(substr)
		if err != nil {
			c.log.Errorf("Conversion error: %v\n", err)
		} else {
			root := math.Sqrt(float64(number))
			rootInt := int(math.Floor(root))
			rootStr := fmt.Sprintf("%d", rootInt)
			var x int
			var y int

			for i := 0; i < len(rootStr); i++ {
				digitStr := string(rootStr[i])
				digit, _ := strconv.Atoi(digitStr)
				if i == 0 {
					x = c.polybiusMapping(digit, "decode")
				} else if i == 1 {
					y = c.polybiusMapping(digit, "decode")
				}
			}

			output += string(c.polybiusSquare[x][y])
		}
	}

	return output
}

func (c *CiphersController) polybiusMapping(index int, operation string) int {
	if operation == "code" {
		switch index {
		case 0:
			return 1
		case 1:
			return 2
		case 2:
			return 3
		case 3:
			return 4
		case 4:
			return 5
		case 5:
			return 6
		case 6:
			return 7
		}
	} else if operation == "decode" {
		switch index {
		case 1:
			return 0
		case 2:
			return 1
		case 3:
			return 2
		case 4:
			return 3
		case 5:
			return 4
		case 6:
			return 5
		case 7:
			return 6
		}
	}

	return -1
}
