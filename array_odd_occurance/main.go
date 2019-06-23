package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    occMap := make(map[int]bool)
    
    for _, inVal := range A {
        if _, present := occMap[inVal]; present {
            delete(occMap, inVal)
        } else {
            occMap[inVal] = true
        }
    }
    
    if len(occMap) == 0 {
       panic("No unpaired values found")
    } else if len(occMap) > 1 {
        panic("Multiple unpaired values found")
    }
    
    var unpairedVal int
    for unpairedVal = range occMap {}
    
    return unpairedVal
}