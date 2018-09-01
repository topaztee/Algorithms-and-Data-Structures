package main

import (
	"fmt"
	"strings"
)

func main(){
    var board = [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},
    }

    board[0][1] = "X"
    board[1][2] = "O"

    for i:= 0; i < len(board); i++ {
        fmt.Println(strings.Join(board[i], " "))
    }
}
var board = [][]string{
}