package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    east_t := 0
    total := 0
    for _, value := range A{
        if value == 0{
            east_t++
        } else if value == 1{
            total += east_t
        }
        if total > 1000000000{
            return -1
        }
    }
    return total
}
