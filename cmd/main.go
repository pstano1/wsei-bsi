package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/pstano1/wsei-bsi/internal/ciphers"
	"github.com/pstano1/wsei-bsi/internal/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var keys utils.Config
	var auxKeys utils.AuxConfig

	err = json.Unmarshal(jsonData, &auxKeys)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	keys.Characters = utils.ConvertStringsToRunes(auxKeys.Characters)
	keys.HomophonicMap = utils.ConvertStringMapToRuneMap(auxKeys.HomophonicMap)
	keys.PolybiusSquare = utils.ConvertStringSlicesToRuneSlices(auxKeys.PolybiusSquare)

	logger := logrus.New()

	ciphersController, err := ciphers.NewCiphersController(logger.WithField("component", "ciphers"), keys.Characters, keys.PolybiusSquare, keys.HomophonicMap)
	if err != nil {
		panic("beale map contains duplicates, aborting")
	}

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
		case "homophonic":
			input := ciphersController.ClearInput(*text)
			result := ciphersController.CodeHomophonic(input)
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
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear, polybius square, homophonic, trithemius & vigenere", *cipher)
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
		case "homophonic":
			result := ciphersController.DecodeHomophonic(*text)
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
			log.Fatalf("Invalid cipher name %s. Supported ciphers: cesear, polybius, homophonic, trithemius & vigenere", *cipher)
		}
	}
}
