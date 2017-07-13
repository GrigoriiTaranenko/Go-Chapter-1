package main

import (
	"awesomeProject/popcount"
	"fmt"
	"strconv"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		y, err := strconv.ParseUint(arg, 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		x := popcount.PopCountr25(y)
		fmt.Printf("x = %d \n", x)
	}
}


