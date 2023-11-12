// The Trithemius cipher is a simple substitution cipher where each letter
// is shifted by a progressively increasing value based on the position in the text.
// This implemenation featrues a key of a single rune that is then convereted to starting point.
package ciphers

func (c *CiphersController) CodeTrithemius(inputString string, key rune) string {
	var output string
	index := c.searchForRune(key, c.characterSet)
	for _, character := range inputString {
		output += c.CodeCaesar(string(character), index)
		index++
	}

	return output
}

func (c *CiphersController) DecodeTrithemius(inputString string, key rune) string {
	var output string
	index := c.searchForRune(key, c.characterSet)
	for _, character := range inputString {
		output += c.DecodeCaesar(string(character), index)
		index++
	}

	return output
}
