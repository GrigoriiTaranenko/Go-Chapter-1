package main

import (
	"os"
	"strconv"
	"fmt"
	"tempconf"
	"distanceconf"
	"weightconf"
)

func commandLine(arg string, units *map[string]float64) {
	conf, err := strconv.ParseFloat(arg, 64)
	if err !=nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconf.Fahrenheit(conf)
	c := tempconf.Celsius(conf)

	m := distanceconf.Meter(conf)
	foot := distanceconf.Foot(conf)

	kg := weightconf.Kg(conf)
	pound := weightconf.Pound(conf)
	fmt.Printf("%s = %s %s = %s \n",
		f, tempconf.FToC(f), c, tempconf.CToF(c))
	fmt.Printf("%s = %s %s = %s \n",
		m, distanceconf.MToF(m), foot, distanceconf.FToM(foot))
	fmt.Printf("%s = %s %s = %s \n",
		kg, weightconf.KToP(kg), pound, weightconf.PToK(pound))
}

func main () {
	units := make(map[string]float64)
	for _, arg := range os.Args[1:] {
		commandLine(arg, &units)
	}
}