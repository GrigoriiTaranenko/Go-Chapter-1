package main

import (
	"os"
	"strconv"
	"fmt"
	"tempconf"
)

func commandLine(arg string, units *map[string]float64) {
	t, err := strconv.ParseFloat(arg, 64)
	if err !=nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconf.Fahrenheit(t)
	c := tempconf.Celsius(t)
	fmt.Printf("%s = %s %s = %s \n",
		f, tempconf.FToC(f), c, tempconf.CToF(c))
}

func main () {
	units := make(map[string]float64)
	for _, arg := range os.Args[1:] {
		commandLine(arg, &units)
	}
}