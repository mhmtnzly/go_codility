package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
    if X >= Y{
        return 0
    }
    distance := Y - X
    count := distance / D
    if distance % D > 0 {
        count++
    }
    return count
}
