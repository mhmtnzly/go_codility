package solution

// you can also use imports, for example:
// import "fmt"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    sort.Ints(A)
    max := A[0] * A[1] * A[2]
    if max < (A[len(A)-1] * A[len(A)-2] * A[len(A)-3]){
        max = A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
    }
    if max < (A[0] * A[1] * A[len(A)-1]){
        max = (A[0] * A[1] * A[len(A)-1])
    }
    // fmt.Println(A, max)
    return max
}