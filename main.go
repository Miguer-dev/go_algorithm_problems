package main

import "fmt"

func main() {
	arr := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(arr))
}
