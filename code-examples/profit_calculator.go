package main

import "fmt"

func profitCalculator() {
	//var revenue, expenses, taxRate float64

	// fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)
	revenue := getUserInput("Revenue: ")

	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)
	expenses := getUserInput("Expenses: ")

	//fmt.Print("Tax Rate: ")
	//fmt.Scan(&taxRate)
	taxRate := getUserInput("Tax Rate: ")

	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate/100)
	// ratio := earningsBeforeTax / profit

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// fmt.Println(earningsBeforeTax)
	// fmt.Println(profit)
	// fmt.Println(ratio)

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f", ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
