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

var logger = logrus.New()

var c = NewCiphersController(logger, characters, polybiusSquare)

func TestInputClearing(t *testing.T) {
	var tests = []struct {
		name    string
		message string
		want    string
	}{
		{"empty message", "", ""},
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
