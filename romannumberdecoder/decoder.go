package romannumberdecoder

import "strings"

func Decode(roman string) int {
	romans := makeRomansNumberMap()
	var decimal int

	lastNumVal := 1
	for _, num := range roman {

		if lastNumVal < romans[strings.ToUpper(string(num))] {
			decimal = romans[strings.ToUpper(string(num))] - lastNumVal
			continue
		}

		decimal += romans[strings.ToUpper(string(num))]
	}

	return decimal
}

func makeRomansNumberMap() map[string]int {
	romans := make(map[string]int)

	romans["M"] = 1000
	romans["D"] = 500
	romans["C"] = 100
	romans["L"] = 50
	romans["X"] = 10
	romans["V"] = 5
	romans["I"] = 1

	return romans
}
