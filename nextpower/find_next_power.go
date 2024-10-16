package nextpower

import "math"

func FindNextPower(val, pow int) int {
	var result float64
	for i := 1; i <= val; i++ {
		result = math.Pow(float64(i), float64(pow))

		if result > float64(val) {
			break
		}
	}

	return int(result)
}
