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

	logger := logrus.New()

	ciphersController := ciphers.NewCiphersController(logger.WithField("component", "ciphers"), characters)

	action := flag.String("action", "", "Action (cipher or decipher)")
	cipher := flag.String("cipher", "", "Cipher name")
	text := flag.String("text", "", "Text to encrypt/decrypt")
	key := flag.String("key", "", "Cipher key")

	flag.Parse()

	if *action != "cipher" && *action != "decipher" {
		log.Fatal("Unsupported action value")
	}
	if *action != "" && (*cipher == "" || *text == "" || *key == "") {
		log.Fatal("Usage: --cipher <cipher_name> --text <text> --key <key> is required")
	}

	if *action == "cipher" {
		switch *cipher {
		case "caesar":
			key, _ := strconv.ParseInt(*key, 10, 64)
			input := strings.ToLower(strings.ReplaceAll(*text, " ", ""))
			result := ciphersController.CipherCaesar(input, int(key))
			fmt.Println("Result:", result)
		default:
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear", *cipher)
		}
	} else if *action == "decipher" {
		switch *cipher {
		case "caesar":
			key, _ := strconv.ParseInt(*key, 10, 64)
			result := ciphersController.DecipherCaesar(*text, int(key))
			fmt.Println("Result:", result)
		default:
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear", *cipher)
		}
	}
}
