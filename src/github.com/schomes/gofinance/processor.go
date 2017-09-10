package gofinance

func CalculateMetrics(transactions []Transaction) Metrics {
	var totalIncome float32
	var totalExpenses float32
	categoryMetrics := make(map[string]float32)
	typeMetrics := make(map[string]float32)

	for _, transaction := range transactions {
		// TODO: change to switch statement?
		if transaction.TransactionType == "credit" {
			totalIncome += transaction.Amount
		} else if transaction.TransactionType == "debit" {
			totalExpenses += transaction.Amount
		}
		categoryMetrics[transaction.Category] += transaction.Amount
		typeMetrics[transaction.Type] += transaction.Amount
	}

	return Metrics{
		TotalIncome: totalIncome,
		TotalExpenses: totalExpenses,
		NetIncome: (totalIncome - totalExpenses),
		CategoryMetrics: categoryMetrics,
		TypeMetrics: typeMetrics,
	}
}