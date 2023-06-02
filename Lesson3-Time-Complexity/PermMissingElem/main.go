package solution

// you can also use imports, for example:
import "sort"
// import "fmt"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // Implement your solution here
    sort.Ints(A)
    if len(A) == 0{
        return 1
    }
    count := 0
    for _, value := range A{
        count ++
        // fmt.Println(count, value)
        if value != count {
            return count
        }
    }  
    return A[len(A)-1]+1
}