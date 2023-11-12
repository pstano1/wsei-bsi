// The Vigenere cipher is a method of encrypting alphabetic text by using a simple form
// of polyalphabetic substitution. The key is repeated to match the length of the plaintext,
// and each letter is shifted according to the corresponding letter in the key.
package ciphers

func (c *CiphersController) CodeVigenere(inputString string, key string) string {
	var output string
	keyList := []rune(key)
	counter := 0
	for _, character := range inputString {
		shift := c.searchForRune(keyList[counter], c.characterSet)
		output += c.CodeCaesar(string(character), shift)
		if counter == (len(keyList) - 1) {
			counter = 0
			continue
		}
		counter++
	}

	return output
}

func (c *CiphersController) DecodeVigenere(inputString string, key string) string {
	var output string
	keyList := []rune(key)
	counter := 0
	for _, character := range inputString {
		shift := c.searchForRune(keyList[counter], c.characterSet)
		output += c.DecodeCaesar(string(character), shift)
		if counter == (len(keyList) - 1) {
			counter = 0
			continue
		}
		counter++
	}

	return output
}
