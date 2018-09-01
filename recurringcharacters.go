package main

import (
	"fmt"
)

func main(){
	str := "abcdea"
	fmt.Println(getFirstRecurringCharacter(str))
}

func getFirstRecurringCharacter(str string) string{
	var hash = make(map[string]bool, len(str))
	
	for _,v := range str {
		if hash[string(v)] == true {
			fmt.Println("duplicate char!", hash[string(v)])
			return string(v)
		}
		hash[string(v)] = true
	}
return "no repeats"
}
