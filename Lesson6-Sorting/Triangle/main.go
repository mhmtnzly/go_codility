package solution

// you can also use imports, for example:
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // Implement your solution here
    if len(A)<3{
        return 0
    }
    sort.Ints(A)
    for i:=1; i<len(A)-1; i++{
        if ((A[i-1]+A[i])>A[i+1])&&((A[i]+A[i+1])>A[i-1])&&((A[i-1]+A[i+1])>A[i]){
            return 1
        }
    }
    return 0