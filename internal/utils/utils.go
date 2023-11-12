package utils

type Config struct {
	Characters     []rune          `json:"characters"`
	PolybiusSquare [][]rune        `json:"polybiusSquare"`
	HomophonicMap  map[rune][]rune `json:"homophonicMap"`
}

type AuxConfig struct {
	Characters     []string            `json:"characters"`
	PolybiusSquare [][]string          `json:"polybiusSquare"`
	HomophonicMap  map[string][]string `json:"homophonicMap"`
}

func ConvertStringsToRunes(strs []string) []rune {
	runes := make([]rune, len(strs))
	for i, str := range strs {
		if len(str) > 0 {
			runes[i] = []rune(str)[0]
		}
	}
	return runes
}

func ConvertStringSlicesToRuneSlices(strSlices [][]string) [][]rune {
	runeSlices := make([][]rune, len(strSlices))
	for i, strSlice := range strSlices {
		runeSlices[i] = ConvertStringsToRunes(strSlice)
	}
	return runeSlices
}

func ConvertStringMapToRuneMap(strMap map[string][]string) map[rune][]rune {
	runeMap := make(map[rune][]rune)
	for key, strSlice := range strMap {
		runeKey := []rune(key)[0]
		runeValue := ConvertStringsToRunes(strSlice)
		runeMap[runeKey] = runeValue
	}
	return runeMap
}
