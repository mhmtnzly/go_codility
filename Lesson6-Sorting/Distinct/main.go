package solution

// you can also use imports, for example:

// import "os"
import "sort"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    if len(A) == 0{
        return 0
    }
    if len(A) == 1 {
        return 1
    }
    sort.Ints(A)
    last_number := A[0]
    count := 1
    for i:=1; i<len(A); i++{
        if last_number != A[i]{
            count +=1
        }
        last_number = A[i]
    }
    return count
}