package ciphers

import (
	"testing"

	"github.com/sirupsen/logrus"
)

var characters = []rune{
	'a', 'ą', 'b', 'c', 'ć', 'd', 'e', 'ę', 'f', 'g', 'h', 'i', 'j', 'k',
	'l', 'ł', 'm', 'n', 'ń', 'o', 'ó', 'p', 'q', 'r', 's', 'ś', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'ź', 'ż',
}

var polybiusSquare = [][]rune{
	{'q', 't', 'ć', 'j', 'y', 'x', 'w'},
	{'ą', 'd', 'e', 'f', 'k', 'ń', 'i'},
	{'ę', 'b', 'v', 'a', 'l', 'o', 'v'},
	{'g', 's', 'c', 'ź', 'n', 'h', 'u'},
	{'ś', 'p', 'ó', 'ł', 'm', 'z', 'ź'},
}

var bealeMap = map[rune][]rune{
	'a': {'e', '2', '3', 'j', '^', '$', '&', '0', ')'},
	'ą': {'f'},
	'b': {'g'},
	'c': {'h', '8', '9', 'm'},
	'ć': {'i'},
	'd': {'j', '0', '!'},
	'e': {'k', '1', '@', 'p', '-', '+', '.', '6'},
	'ę': {'l'},
	'f': {'m'},
	'g': {'n'},
	'h': {'o'},
	'i': {'p', '6', '&', 'u', '(', ')', '-', '1'},
	'j': {'q', '7'},
	'k': {'r', '8', '(', 'w'},
	'l': {'s', '9'},
	'ł': {'t', '0'},
	'm': {'u', '!', '+'},
	'n': {'v', '@', ',', 'ź', ',', '<'},
	'ń': {'o'},
	'o': {'x', '=', '|', ';', '9', ';'},
	'ó': {'y'},
	'p': {'z', '^', '\\'},
	'q': {'ź'},
	'r': {'ż', '*', '?', 'ć', '?'},
	's': {'d', '(', '<', 'd'},
	'ś': {'f'},
	't': {'g', '*', '?', 'ę'},
	'u': {'h', '+', '!'},
	'v': {'i'},
	'w': {'j', '-', '#', 'h', '*', '('},
	'x': {'k'},
	'y': {'l', '/', '%', 'j'},
	'z': {'m', '0', '^', 'k', '*', '+'},
	'ź': {'n'},
	'ż': {'o', '2', '`', 'ł'},
}

var logger = logrus.New()

var c = NewCiphersController(logger, characters, polybiusSquare, bealeMap)

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

func TestBealeCoding(t *testing.T) {
	var tests = []struct {
		name    string
		message string
	}{
		{"empty message", ""},
		{"word: test", "test"},
		{"whole charset", "aąbcćdeęfghijklłmnńoópqrsśtuvwxyzźż"},
	}

	for _, test := range tests {
		output := c.CodeBeale(test.message)
		if len(output) != len(test.message) {
			t.Errorf("Input: %s - Output: %s - Expected length: %d - Got: %d", test.message, output, len(test.message), len(output))
		}
		if len(output) == 0 && len(test.message) != 0 {
			t.Errorf("Input: %s - Output is empty", test.message)
		}
	}
}
