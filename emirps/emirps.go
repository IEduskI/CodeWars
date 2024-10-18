package emirps

import "strconv"

func FindEmirp(n int) [3]int {
	var emirps, lastEmirp, largestEmirp, sumEmirps int
	// iterate all numbers before n
	for num := 1; num <= n; num++ {

		if isPrime(num) {
			newInt := reverseNumber(num)

			if num == newInt {
				continue
			}

			if isPrime(newInt) {
				emirps++

				if num > lastEmirp {
					largestEmirp = num
				}

				sumEmirps += num
				lastEmirp = num
			}
		}
	}

	return [3]int{emirps, largestEmirp, sumEmirps}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverseNumber(n int) int {
	intString := strconv.Itoa(n)

	runes := []rune(intString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	newInt, _ := strconv.Atoi(string(runes))

	return newInt
}
