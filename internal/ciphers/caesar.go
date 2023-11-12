// Caesar cipher is a substitution cipher where each letter in the plaintext
// is shifted a certain number of places down or up the alphabet.
package ciphers

func (c *CiphersController) CodeCaesar(inputString string, shift int) string {
	var output string

	for _, character := range inputString {
		index := c.searchForRune(character, c.characterSet)

		index += shift
		if index > 34 {
			index -= 35
		}
		output += string(c.characterSet[index])
	}

	return output
}

func (c *CiphersController) DecodeCaesar(inputString string, shift int) string {
	var output string

	for _, character := range inputString {
		index := c.searchForRune(character, c.characterSet)

		index -= shift
		if index < 0 {
			index += 35
		}
		output += string(c.characterSet[index])
	}

	return output
}
