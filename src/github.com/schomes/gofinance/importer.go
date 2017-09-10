package gofinance

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"os"
)

/*
 * Assumptions:
 * While the order of columns doesn't matter, the names do. A transaction consists of the following; and
 * the columns must be named exactly:
 * - Date
 * - Description
 * - Amount
 * - Transaction Type
 * - Category
 * - Type
 *
 * Extra columns are allowed, but will be ignored.
 */
func ImportFromCSV(fileName string) ([]Transaction, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	line, err := reader.Read()
	if err == io.EOF {
		return nil, fmt.Errorf("file is empty")
	} else if err != nil {
		return nil, fmt.Errorf("cannot parse file: %v", err)
	}

	// Associate each column number (index) to the title of the column. This allows
	// columns to be listed in any order in a csv file.
	columnNameIndicies := make(map[string]int)
	for index, columnName := range line {
		columnNameIndicies[columnName] = index
	}

	// TODO: if the column name doesn't exist (such as if Type isn't present) and we try to
	// get the index for this column, 0 (default value) is returned. This defaults to getting
	// the value of the first column (this will probably be undesired behavior)

	var transactions []Transaction
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("cannot parse file: %v", err)
		}

		// Getting the Amount from the csv file is a little more involved, because it requires
		// converting a string into a float32. Therefore, we do some work outside of the append function below.
		amount, err := strconv.ParseFloat(line[columnNameIndicies["Amount"]], 32)
		if err != nil {
			return nil, fmt.Errorf("cannot read 'Amount' column %v", err)
		}

		// strconv.ParseFloat returns a float64 even when 32 is the specified bit size; we get a float that
		// is convertible to float32.
		// see: https://golang.org/pkg/strconv/#ParseFloat
		transactions = append(transactions, Transaction{
			Date: line[columnNameIndicies["Date"]],
			Description: line[columnNameIndicies["Description"]],
			Amount: float32(amount),
			TransactionType: line[columnNameIndicies["Transaction Type"]],
			Category: line[columnNameIndicies["Category"]],
			Type: line[columnNameIndicies["Type"]],
		})
	}

	return transactions, nil
}