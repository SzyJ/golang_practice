package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	smallest := 1

	laMap := make(map[int]bool)

	for _, arrVal := range A {

		if smallest == arrVal {
			smallest += 1

			for laMap[smallest] {
				delete(laMap, smallest)
				smallest += 1
			}

		} else if smallest < arrVal {
			laMap[arrVal] = true
		}
	}

	return smallest
}
