package solution

// you can also use imports, for example:
// import "fmt"
import "strings"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, P []int, Q []int) []int {
    m := len(P)
    result := make([]int, m)

    for i := 0; i < m; i++ {
        start := P[i]
        end := Q[i] + 1
        slice := S[start:end]

        if strings.Contains(slice, "A") {
            result[i] = 1
            continue
        } else if strings.Contains(slice, "C") {
            result[i] = 2
        } else if strings.Contains(slice, "G") {
            result[i] = 3
        } else {
            result[i] = 4
        }
    }

    return result
}