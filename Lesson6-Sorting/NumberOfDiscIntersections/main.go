package solution
// Complexity is high ????????????????????????
// O(N*log(N)) or O(N) or O(N**2)
// you can also use imports, for example:
// import "fmt"
// import "sort"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    count := 0
    for index, value := range A{
        for i := index-1; i>=0; i--{
            if index-value<=i-A[i] || A[i]+i>=index-value{
                count++
                if count > 10000000{
                    return -1
            }
            }
            
        }
    }
    return count
}