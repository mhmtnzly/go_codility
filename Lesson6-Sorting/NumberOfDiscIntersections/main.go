package solution

// you can also use imports, for example:
// import "fmt"
// import "sort"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // sort.Ints(A)\
    // length := len(A)
    count := 0
    for index, value := range A{
        // fmt.Println(index-value, index+value)
        for i := index-1; i>=0; i--{
            if index-value<=i-A[i] || A[i]+i>=index-value{
                count++
                // fmt.Println(i-A[i],index-value, value, index, i)
            }
            if count > 10000000{
            return -1
        }
        }
        
    }
    return count
}