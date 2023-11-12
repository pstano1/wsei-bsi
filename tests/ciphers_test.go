package ciphers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pstano1/wsei-bsi/internal/ciphers"
	"github.com/pstano1/wsei-bsi/internal/utils"
	"github.com/sirupsen/logrus"
)

var (
	c      ciphers.ICiphersController
	logger *logrus.Logger
)

func TestMain(m *testing.M) {
	file, err := os.Open("../config.json")
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

	var auxKeys utils.AuxConfig

	err = json.Unmarshal(jsonData, &auxKeys)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	characters := utils.ConvertStringsToRunes(auxKeys.Characters)
	polybiusSquare := utils.ConvertStringSlicesToRuneSlices(auxKeys.PolybiusSquare)
	homophonicMap := utils.ConvertStringMapToRuneMap(auxKeys.HomophonicMap)

	c, _ = ciphers.NewCiphersController(logger, characters, polybiusSquare, homophonicMap)

	m.Run()
}

func TestInputClearing(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		want    string
	}{
		{"empty message", "", ""},
		{"3x white space", "   ", ""},
		{"white spaces", "abc abc", "abcabc"},
		{"punctuation", "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", ""},
		{"lowercase", "ABC abc", "abcabc"},
		{"whole paragraph", "- Zło to zło, Stregoborze - rzekł poważnie wiedźmin wstając. - Mniejsze, większe, średnie, wszystko jedno, proporcje są umowne a granice zatarte. Nie jestem świątobliwym pustelnikiem, nie samo dobro czyniłem w życiu. Ale jeżeli mam wybierać pomiędzy jednym złem a drugim, to wolę nie wybierać wcale.", "złotozłostregoborzerzekłpoważniewiedźminwstającmniejszewięksześredniewszystkojednoproporcjesąumowneagranicezatarteniejestemświątobliwympustelnikiemniesamodobroczyniłemwżyciualejeżelimamwybieraćpomiędzyjednymzłemadrugimtowolęniewybieraćwcale"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := c.ClearInput(test.message)
			if res != test.want {
				t.Errorf("got %s, want %s", res, test.want)
			}
		})
	}
}

func TestCaesarEncoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		key     int
		want    string
	}{
		{"empty message", "", 5, ""},
		{"negative key", "źżxy", -3, "xyuv"},
		{"key of 0", "ąćę", 0, "ąćę"},
		{"whole charset by 3", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż", 3, "cćdeęfghijklłmnńoópqrsśtuvwxyzźżaąb"},
		{"whole charset by 35", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż", 35, "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := c.CodeCaesar(test.message, test.key)
			if res != test.want {
				t.Errorf("got %s, want %s", res, test.want)
			}
		})
	}
}

func TestCaesarDecoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		key     int
		want    string
	}{
		{"empty message", "", 5, ""},
		{"negative key", "xyuv", -3, "źżxy"},
		{"key of 0", "ąćę", 0, "ąćę"},
		{"whole charset by 3", "cćdeęfghijklłmnńoópqrsśtuvwxyzźżaąb", 3, "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
		{"whole charset by 35", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż", 35, "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := c.DecodeCaesar(test.message, test.key)
			if res != test.want {
				t.Errorf("got %s, want %s", res, test.want)
			}
		})
	}
}

func TestPolybiusCoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		want    string
	}{
		{"empty message", "", ""},
		{"message: test", "test", "144 529 1764 144 "},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := c.CodePolybiusSquare(test.message)
			if res != test.want {
				t.Errorf("got %s, want %s", res, test.want)
			}
		})
	}
}

func TestPolybiusDecoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		want    string
	}{
		{"empty message", "", ""},
		{"message: test", "144 529 1764 144", "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := c.DecodePolybiusSquare(test.message)
			if res != test.want {
				t.Errorf("got %s, want %s", res, test.want)
			}
		})
	}
}

func TestHomophonicCoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
	}{
		{"empty message", ""},
		{"word: test", "test"},
		{"whole charset", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := c.CodeHomophonic(test.message)
			runes := []rune(output)
			originalRunes := []rune(test.message)
			if len(runes) != len(originalRunes) {
				t.Errorf("Input: %s - Output: %s - Expected length: %d - Got: %d", test.message, output, len(originalRunes), len(runes))
			}
			if len(runes) == 0 && len(originalRunes) != 0 {
				t.Errorf("Input: %s - Output is empty", test.message)
			}
		})
	}
}

func TestHomophonicDecoding(t *testing.T) {
	var tests = []struct {
		name           string
		message        string
		expectedOutput string
	}{
		{"empty message", "", ""},
		{"word: test", "☼☀_☼", "test"},
		{"word test #2", "☼☀_★", "test"},
		{"whole charset", ":f§mi1☀>noqr♦}u'ńóż#~\\*śzy☼4¶d)25%6", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
		{"whole charset #2", "0f§äi.☀>noq7♦tu'ń;ż#~♥*ś_$★!¶d)2v%6", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := c.DecodeHomophonic(test.message)
			if output != test.expectedOutput {
				t.Errorf("got %s, want: %s", output, test.expectedOutput)
			}
		})
	}
}

func TestTrithemiusCoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		key     rune
		want    string
	}{
		{"empty message", "", 'a', ""},
		{"key of \"A\" => first shift is 0", "abc", 'a', "acd"},
		{"non-empty message with first shift of 3", "hello", 'c', "khoót"},
		{"message with diacritics with first shift of 5", "ąbcć", 'd', "efhj"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := c.CodeTrithemius(test.message, test.key)
			if output != test.want {
				t.Errorf("got %s, want %s", output, test.want)
			}
		})
	}
}

func TestTrithemiusDecoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		key     rune
		want    string
	}{
		{"empty message", "", 'a', ""},
		{"key of \"A\" => first shift is 0", "acd", 'a', "abc"},
		{"non-empty message with first shift of 3", "khoót", 'c', "hello"},
		{"message with diacritics with first shift of 5", "efhj", 'd', "ąbcć"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := c.DecodeTrithemius(test.message, test.key)
			if output != test.want {
				t.Errorf("got %s, want %s", output, test.want)
			}
		})
	}
}

func TestVigenereCoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		key     string
		want    string
	}{
		{"empty message", "", "key", ""},
		{"non-empty message", "hello", "abc", "hfnlp"},
		{"message with diacritics", "ąbcć", "def", "ąijł"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := c.CodeVigenere(test.message, test.key)
			if output != test.want {
				t.Errorf("got %s, want %s", output, test.want)
			}
		})
	}
}

func TestVigenereDecoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		key     string
		want    string
	}{
		{"empty message", "", "key", ""},
		{"non-empty message", "hfnlp", "abc", "hello"},
		{"message with diacritics", "ąijł", "def", "ąbcć"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := c.DecodeVigenere(test.message, test.key)
			if output != test.want {
				t.Errorf("got %s, want %s", output, test.want)
			}
		})
	}
}
