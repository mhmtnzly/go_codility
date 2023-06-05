package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
    list := make([]int, N)
    max_counter := 0
    add_last := 0
    for _, value := range A{
        if value>=1 && value<=N{
            list[value-1] += 1
            if max_counter<list[value-1]{
                max_counter = list[value-1]
            }
        }else{
            add_last += max_counter
            list = make([]int, N)
        }
    }
    for index, _ := range list{
        list[index] += add_last
    }
    return list
}
