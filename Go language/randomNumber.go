package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	//You’ll always see the same sequence every time you run the program. The random number changes inside the program,
	//you need to provide a unique seed for your program. You really want to not forget seeding, and instead properly seed our pseudonumber generator. How?

//	Use rand.Seed() before calling any math/rand method, passing an int64 value. You just need to seed once in your program, not every time you need a random number. The most used seed is the current time, converted to int64 by UnixNano with rand.Seed(time.Now().UnixNano()):
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(5))
    fmt.Println(rand.Intn(5))
    fmt.Println(rand.Intn(5))
    fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	fmt.Println()
	fmt.Println(randomInt(1, 11)) //get an int in the 1...10 range
	fmt.Println(randomArray(10))
}


// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}


func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}
