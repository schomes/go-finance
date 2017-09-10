// TODO: doc comment goes here
package gofinance

//             _.-'~~`~~'-._
//         .'`             `'.
//        /                    \
//       /`       .-'~"-.       `\
//      ;        / `-    \        ;
//     ;        />  `.  -.|        ;
//     |       /_     '-.__)       |
//     |        |-  _.' \ |        |
//     ;        `~~;     \\        ;
//      ;  GO      /      \\)P    ;
//       \ FINANCE '.___.-'`"    /
//        `\                   /`
//          '._   2 0 1 7   _.'
//      jgs    `'-..,,,..-'`

import (
	"fmt"
	"time"
)

type Transaction struct {
	Date            string
	Description     string
	Amount          float32
	TransactionType string
	Category        string
	Type            string
}

type Metrics struct {
	TotalIncome float32
	TotalExpenses float32
	NetIncome float32
	CategoryMetrics map[string]float32
	TypeMetrics map[string]float32
}

// TODO: messy, quick function; should be redone
// Returns the download URL for a Mint CSV file.
// The file will contain transactions between the
// startDate and endDate.
func GetMintURL(startDate time.Time, endDate time.Time) string {
	startYear, startMonth, startDay := startDate.Date()
	endYear, endMonth, endDay := endDate.Date()
	spaceChar := "%%2F"

	startMonthInt := int(startMonth)
	endMonthInt := int(endMonth)

	startDateString := fmt.Sprintf("%d%s%d%s%d", startMonthInt, spaceChar, startDay, spaceChar, startYear)
	endDateString := fmt.Sprintf("%d%s%d%s%d", endMonthInt, spaceChar, endDay, spaceChar, endYear)

	url := fmt.Sprintf("https://mint.intuit.com/transactionDownload.event?startDate=%s&endDate=%s&offset=0&filterType=cash&comparableType=8", startDateString, endDateString)
	return url
}

// TODO: consider renaming Type to BudgetClass, BudgetCategory, BudgetType, etc.
