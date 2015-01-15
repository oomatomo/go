package main

import (
	"fmt"
)

func main() {
	id := "hoge"
	hogeID := "hoge"
	fmt.Println("hello world")
	fmt.Println(id)
	fmt.Println(hogeID)
	fmt.Println(Square(2))
}

// Square return the
func Square(n int) int {
	return n * n
}
