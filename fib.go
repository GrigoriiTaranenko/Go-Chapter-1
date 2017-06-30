package main

import (
	"os"
	"fmt"
	"strconv"
)
/**

 */
func fib (n int) int  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}

func main() {
	var x int
	var s string
	s =os.Args[1]
	println(s)
	d,err := strconv.Atoi(s)
	if err!=nil {
		fmt.Println("Error \n", err)

	}
	x = fib(d)
	fmt.Printf("Пример %d \n", x)
}