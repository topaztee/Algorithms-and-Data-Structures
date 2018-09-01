package main

import "fmt"

func f(s [s]string, level int) {
        if level > 5 {
               return
        }
        s = append(s, fmt.Sprint(level))
        f(s, level+1)
        fmt.Println("level:", level, "slice:", s)
}

func main() {
        f(nil, 0)
}

//You can see that at each level the value of s was unaffected by the operation 
//of other callers of f, and that while four underlying arrays were created
// 6 higher levels of f in the call stack are unaffected by the copy and 
//reallocation of new underlying arrays as a by-product of append.

