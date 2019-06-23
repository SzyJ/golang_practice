package solution

import (
	"strconv"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	gapCounter := 0
	largestGap := 0

	binString := getBinary(N)

	for _, char := range binString {
		if char == '1' {
			if largestGap < gapCounter {
				largestGap = gapCounter
			}

			gapCounter = 0
		} else if char == '0' {
			gapCounter += 1
		} else {
			panic("invalid binary string: " + binString)
		}
	}

	return largestGap
}

func getBinary(N int) string {
	return strconv.FormatInt(int64(N), 2)
}
