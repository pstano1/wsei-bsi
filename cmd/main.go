package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/pstano1/wsei-bsi/internal/ciphers"
	"github.com/sirupsen/logrus"
)

func main() {
	characters := []rune{
		'a', 'ą', 'b', 'c', 'ć', 'd', 'e', 'ę', 'f', 'g', 'h', 'i', 'j', 'k',
		'l', 'ł', 'm', 'n', 'ń', 'o', 'ó', 'p', 'q', 'r', 's', 'ś', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'ź', 'ż',
	}

	polybiusSquare := [][]rune{
		{'q', 't', 'ć', 'j', 'y', 'x', 'w'},
		{'ą', 'd', 'e', 'f', 'k', 'ń', 'i'},
		{'ę', 'b', 'v', 'a', 'l', 'o', 'r'},
		{'g', 's', 'c', 'ź', 'n', 'h', 'u'},
		{'ś', 'p', 'ó', 'ł', 'm', 'z', 'ż'},
	}

	bealeMap := map[rune][]rune{
		'a': {'e', '❤', '♫', 'x', ':', '&', '0', '☺'},
		'ą': {'f'},
		'b': {'§'},
		'c': {'h', '8', '9', 'm'},
		'ć': {'i'},
		'd': {'k', '1', '@', 'p', '-', '!', '.'},
		'e': {'☀'},
		'ę': {'>'},
		'f': {'n'},
		'g': {'o'},
		'h': {'q'},
		'i': {'r', '7'},
		'j': {'s', '♦'},
		'k': {'t', '}'},
		'l': {'u', '^'},
		'ł': {'\'', '{', ',', 'ź', '<'},
		'm': {'ń'},
		'n': {'ó', '=', '|', ';', '8'},
		'ń': {'ż'},
		'o': {'#'},
		'ó': {'~'},
		'p': {'♣', '\\', '♥'},
		'q': {'*'},
		'r': {'ś', '?'},
		's': {'z', '_'},
		'ś': {'$', 'y'},
		't': {'★', '☼'},
		'u': {'4', '+', '!'},
		'v': {'¶'},
		'w': {'/', 'd'},
		'x': {')'},
		'y': {'¢', '2', '©', 'ę'},
		'z': {'g', '5', 'v', '(', 'f'},
		'ź': {'%'},
		'ż': {'6'},
	}

	logger := logrus.New()

	ciphersController := ciphers.NewCiphersController(logger.WithField("component", "ciphers"), characters, polybiusSquare, bealeMap)

	action := flag.String("action", "", "Action (code or decode)")
	cipher := flag.String("cipher", "", "Cipher name")
	text := flag.String("text", "", "Text to encrypt/decrypt")
	key := flag.String("key", "", "Cipher key")

	flag.Parse()

	if *action != "code" && *action != "decode" {
		log.Fatal("Unsupported action value")
	}
	if *action != "" && (*cipher == "" || *text == "") {
		log.Fatal("--cipher <cipher_name> --text <text> are required")
	}

	if *action == "code" {
		switch *cipher {
		case "caesar":
			if *key == "" {
				log.Fatal("--key <key> is required for this cipher")
			}
			key, _ := strconv.ParseInt(*key, 10, 64)
			input := ciphersController.ClearInput(*text)
			result := ciphersController.CodeCaesar(input, int(key))
			fmt.Println("Result:", result)
		case "polybius":
			input := ciphersController.ClearInput(*text)
			result := ciphersController.CodePolybiusSquare(input)
			fmt.Println("Result:", result)
		case "beale":
			input := ciphersController.ClearInput(*text)
			result := ciphersController.CodeBeale(input)
			fmt.Println("Result:", result)
		case "trithemius":
			key := ciphersController.ClearInput(*key)
			if key == "" {
				log.Fatal("--key <key> is required for this cipher")
			}
			keyRune := []rune(key)
			if len(keyRune) != 1 {
				log.Fatal("Please provide a valid key")
			}
			input := ciphersController.ClearInput(*text)
			result := ciphersController.CodeTrithemius(input, keyRune[0])
			fmt.Println("Result:", result)
		case "vigenere":
			key := ciphersController.ClearInput(*key)
			if key == "" {
				log.Fatal("--key <key> is required for this cipher")
			}
			input := ciphersController.ClearInput(*text)
			result := ciphersController.CodeVigenere(input, key)
			fmt.Println("Result:", result)
		default:
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear, polybius square, beale, trithemius & vigenere", *cipher)
		}
	} else if *action == "decode" {
		switch *cipher {
		case "caesar":
			if *key == "" {
				log.Fatal("--key <key> is required for this cipher")
			}
			key, _ := strconv.ParseInt(*key, 10, 64)
			result := ciphersController.DecodeCaesar(*text, int(key))
			fmt.Println("Result:", result)
		case "polybius":
			result := ciphersController.DecodePolybiusSquare(*text)
			fmt.Println("Result:", result)
		case "beale":
			result := ciphersController.DecodeBeale(*text)
			fmt.Println("Result:", result)
		case "trithemius":
			key := ciphersController.ClearInput(*key)
			if key == "" {
				log.Fatal("--key <key> is required for this cipher")
			}
			keyRune := []rune(key)
			input := ciphersController.ClearInput(*text)
			result := ciphersController.DecodeTrithemius(input, keyRune[0])
			fmt.Println("Result:", result)
		case "vigenere":
			key := ciphersController.ClearInput(*key)
			if key == "" {
				log.Fatal("--key <key> is required for this cipher")
			}
			input := ciphersController.ClearInput(*text)
			result := ciphersController.DecodeVigenere(input, key)
			fmt.Println("Result:", result)
		default:
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear, polybius, beale, trithemius & vigenere", *cipher)
		}
	}
}
