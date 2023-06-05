package solution

// you can also use imports, for example:
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    down := 0 
    total := 0
    list := []int{}
    length := len(A)
    if length == 0 {
        return 0
    }
    for _, value := range A{
        total += value
    }
    for index, value := range A {
        if index+1 == length {
            continue
        }
        down += value
        total -= value
        diff := down - total
        if diff < 0{
            list = append(list, diff*-1)
        }else{
            list = append(list, diff)
        }
    }
    min := list[0]
    for _, value := range list {
        if value < min {
            min = value
        }
    }
    return min
}