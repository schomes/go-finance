package main

import (
	"github.com/schomes/gofinance"
	"os"
	// "flag"
	"fmt"
	// "time"
)

func main() {

	// The CSV filename is the first argument
	if len(os.Args[1:]) == 0 {
		fmt.Printf("Usage: %s filepath\n", os.Args[0])
		return
	}

	fileName := os.Args[1]
	transactions, err := gofinance.ImportFromCSV(fileName)
	if err != nil {
		fmt.Printf("cannot parse file: %v", err)
		return
	}

	metrics := gofinance.CalculateMetrics(transactions)

	gofinance.PrintToTerminal(metrics)

	// TODO: mint URL generator (call if flag given with options (date range))
	/*
	start := time.Date(2017, 9, 1, 0, 0, 0, 0, time.UTC) 
	end := time.Date(2017, 9, 30, 0, 0, 0, 0, time.UTC) 
	fmt.Printf("%s\n", gofinance.GetMintURL(start, end))
	*/
}