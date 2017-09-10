package gofinance

import (
	"fmt"
)

func PrintToTerminal(metrics Metrics) {
	fmt.Printf("==== REPORT ====\n")
	fmt.Printf("Total Income: $%.2f\nTotal Expenses: $%.2f\nNet Income: $%.2f\n", metrics.TotalIncome, metrics.TotalExpenses, metrics.NetIncome)
	fmt.Printf("\n")

	fmt.Printf("~~~~ Categories ~~~~\n")
	for category, amount := range metrics.CategoryMetrics {
		fmt.Printf("%s: %.2f\n", category, amount)
	}
	fmt.Printf("\n")

	fmt.Printf("~~~~ Here's how you're doing for each budget type ~~~~\n")
	for budgetType, amount := range metrics.TypeMetrics {
		// don't print if budget-type name is empty
		if len(budgetType) == 0 {
			continue
		}
		percentTotalIncome := int((amount / metrics.TotalIncome) * 100)
		fmt.Printf("%s: $%.2f (%d%% of total income)\n", budgetType, amount, percentTotalIncome)
	}
}

// func PrintToLatex