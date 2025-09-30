package main

import (
	"fmt" // this is standart library which installed as default
	"math"
)

const inflationRate = 2.5

func investmentCalculator() {
	// fmt.Print("Hello World!")
	// fmt.Print(`Hello World!`) // another using but not common
	// := değişkenin tipini kendisi ayarlar, ekstra var kullanmamıza gerek yok
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 -> tek satırda kullanım
	var investmentAmount, years, expectedReturnRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFuturuvalues(investmentAmount, expectedReturnRate, years)

	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRealFutureValue := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// outputs information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue, futureRealValue)
	// fmt.Print(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)

	fmt.Print(formattedFutureValue, formattedRealFutureValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFuturuvalues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	//# Method-1:
	// fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// rfv := fv / math.Pow(1+inflationRate/100, years)

	//# Method-2
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return // (fv float64, rfv float64) yazdığımızda direkt olarak return kullanmamız yeterli.
}
