package solution

// you can also use imports, for example:
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    length := len(A)
    list := make([]int, length)
    for _, value := range A{

        if value > length || value == 0 {
            return 0
        }
        if list[value-1] == 1{
            return 0
        } else{
            list[value-1] = 1
        }
    }
    for _, value := range list{
        if value == 0{
            return 0
        }
    }
    return 1
}