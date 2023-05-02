package solution

func Solution(A []int, K int) []int {
    if len(A) == 0 {
        return A
    }else if K == 0 {
        return A
    }
    if K>len(A){
        K = K%len(A)
    }
    var last int
    for i:=0; i<K; i++{
        last = A[len(A)-1]
        A = append([]int{last}, A...)
        A = A[:len(A)-1]
    }
    return A
}