package ciphers

func (c *CiphersController) CipherCaesar(inputString string, shift int) string {
	var output string

	for _, character := range inputString {
		index := c.searchForRune(character)

		index += shift
		if index > 34 {
			index -= 35
		}
		output += string(c.characterSet[index])
	}

	return output
}

func (c *CiphersController) DecipherCaesar(inputString string, shift int) string {
	var output string

	for _, character := range inputString {
		index := c.searchForRune(character)

		index -= shift
		if index > 34 {
			index -= 35
		}
		output += string(c.characterSet[index])
	}

	return output
}
