package solution

// you can also use imports, for example:
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    count := 1
    length := len(A)
    sort.Ints(A)
    for index, value := range A{
        if count == value {
            count++
            if index == length-1{
                return value + 1
			}
        }
    }
    if count > 1 {
        return count
    }
    return 1
}