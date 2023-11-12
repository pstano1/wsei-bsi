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
