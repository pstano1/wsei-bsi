package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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
		{'ę', 'b', 'v', 'a', 'l', 'o', 'v'},
		{'g', 's', 'c', 'ź', 'n', 'h', 'u'},
		{'ś', 'p', 'ó', 'ł', 'm', 'z', 'ź'},
	}

	logger := logrus.New()

	ciphersController := ciphers.NewCiphersController(logger.WithField("component", "ciphers"), characters, polybiusSquare)

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
			input := strings.ToLower(strings.ReplaceAll(*text, " ", ""))
			result := ciphersController.CodeCaesar(input, int(key))
			fmt.Println("Result:", result)
		case "polybius":
			input := strings.ToLower(strings.ReplaceAll(*text, " ", ""))
			result := ciphersController.CodePolybiusSquare(input)
			fmt.Println("Result:", result)
		default:
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear, polybius", *cipher)
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
			input := strings.ToLower(*text)
			result := ciphersController.DecodePolybiusSquare(input)
			fmt.Println("Result:", result)
		default:
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear, polybius", *cipher)
		}
	}
}
