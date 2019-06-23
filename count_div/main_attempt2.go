package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
    // write your code in Go 1.4
    var firstHit int
    noDivides := true

    for i := A; i <= B; i += 1 {
        if i % K == 0 {
            firstHit = i
            noDivides = false
            break
        }
    }

    if noDivides {
        return 0
    }
    
    diff := B - firstHit
    var occ int = diff / K 
    
    return occ + 1
}