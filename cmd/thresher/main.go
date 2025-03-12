package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var threshold string
	var value string
	flag.StringVar(&threshold, "threshold", "80", "threshold to use")
	flag.StringVar(&value, "value", "", "file to load the data from")
	flag.Parse()
	numThreshold, err := strconv.ParseFloat(threshold, 64)
	if err != nil {
		iNumThreshold, err := strconv.Atoi(threshold)
		if err != nil {
			fmt.Printf("could not parse threshold: %s\n", threshold)
			os.Exit(1)
		}
		numThreshold = float64(iNumThreshold)
	}
	numValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		iNumValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("could not parse value: %s\n", value)
			os.Exit(2)
		}
		numValue = float64(iNumValue)
	}
	if numValue > numThreshold {
		fmt.Printf("value %.2f is greater than threshold %.2f\n", numValue, numThreshold)
		os.Exit(0)
	} else {
		fmt.Printf("value %.2f is less than threshold %.2f\n", numValue, numThreshold)
		os.Exit(5)
	}
}
