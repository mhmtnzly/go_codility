package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
    // Implement your solution here
    list := make([]int, X)
    count := 0
    count_steps := 0
    for _, value := range A{
        if value <= X && list[value-1] != 1{
            list[value-1] = 1
            count += 1
            if count == X {
                return count_steps
            }
        }
        count_steps += 1
    }
    return -1
}
