package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
    counters := make([]int, N)
    maxCounter := 0
    currentMax := 0
    
    for _, value := range A {
        if value >= 1 && value <= N {
            if counters[value-1] < maxCounter {
                counters[value-1] = maxCounter
            }
            counters[value-1]++
            if counters[value-1] > currentMax {
                currentMax = counters[value-1]
            }
        } else if value == N+1 {
            maxCounter = currentMax
        }
    }
    
    for i := 0; i < N; i++ {
        if counters[i] < maxCounter {
            counters[i] = maxCounter
        }
    }
    
    return counters
}