package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
    count := 0
    if A%K == 0 {
        count++
    }
    count += B/K - A/K

    return count
}