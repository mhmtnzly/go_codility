package solution
func Solution(A []int) int {
    dict := make(map[int]int)    
    for _, int := range A {
        dict[int]++
    }
    for key, val := range dict {
        if val%2!=0{
            return key
        }
    }
    return 0
}