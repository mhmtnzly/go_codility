package solution

// you can also use imports, for example:
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
    // Implement your solution here
    length := len(S)
    list := []string{}
    elements := []string{"{","}","[","]", "(", ")"}
    if len(S)==0{
        return 1
    }
    for _, value := range S{
        length = len(list)
        if length>0 {
            if list[length-1]==string(elements[0]) && string(value) == string(elements[1]){
                list = list[:length-1]
                continue
            }else if list[length-1]==string(elements[2]) && string(value) == string(elements[3]){
                list = list[:length-1]
                continue
            }else if list[length-1]==string(elements[4]) && string(value) == string(elements[5]){
                list = list[:length-1]
                continue   
            }
        }
        list = append(list, string(value))
    }
    if len(list)>0{
        return 0
    }
    return 1
}
