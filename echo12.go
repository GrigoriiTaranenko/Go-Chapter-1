package main

import (
	"os"
	"fmt"
	"strconv"
)

func main()  {
	var s, newLine, space  string
	for i := 0; i < len(os.Args); i++ {
		newLine = "\n"
		space = " "
		s += strconv.Itoa(i) + space + os.Args[i] + newLine
	}
	fmt.Println(s)
}