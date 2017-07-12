package main

import (
	"awesomeProject/popcount"
	"fmt"
	"strconv"
	"os"
)

func main() {
	x := make(map[string]int64)
	for _, arg := range os.Args[1:] {
		y, err := strconv.ParseInt(arg, 64, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		x := popcount.PopCount(y)
		fmt.Printf("x = %d \n", x)
	}

}

