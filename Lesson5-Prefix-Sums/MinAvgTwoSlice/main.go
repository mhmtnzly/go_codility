package solution

// you can also use imports, for example:
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    min_index := 0
    min := 10000.0
    for i := 0; i<len(A)-1; i++{
        avg := (float64(A[i]) + float64(A[i+1]))/2
        if min > avg{
            min_index = i
            min = avg
        }
    }
    for i := 0; i<len(A)-2; i++ {
        avg := (float64(A[i]) + float64(A[i+1])+ float64(A[i+2]))/3 
        if min > avg{
            min_index = i
            min = avg 
        }
    }
    return min_index
}