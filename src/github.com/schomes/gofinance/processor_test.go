package gofinance

import (
	"testing"
	"reflect"
)

// TODO: these tests haven't been updated, so they are broken :'(
func TestCalculateMetrics(t *testing.T) {
	var transactions = []Transaction {
		{"6/10/17", "Bob's Foods", 55.24, "debit", "Groceries", "fixed-costs"},
		{"6/11/17", "BrewCoffeeshop", 4.25, "debit", "Coffee Shops", "guilt-free"},
		{"6/13/17", "BrewCoffeeshop", 3.50, "debit", "Coffee Shops", "guilt-free"},
		{"6/15/17", "Direct deposit", 7000.00, "credit", "Paycheck", ""},
		{"6/30/17", "Bus pass", 65.50, "debit", "Public Transportation", "fixed-costs"},
	}
	expectedMetrics := Metrics{TotalIncome: 7000.00, TotalExpenses: 128.49, NetIncome: 6871.51}

	if got := CalculateMetrics(transactions); !reflect.DeepEqual(got, expectedMetrics) {
		t.Errorf("CalculateMetrics(%v): got %v, expected %v", transactions, got, expectedMetrics)
	}
}