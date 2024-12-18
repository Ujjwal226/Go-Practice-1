package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, err = getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
	}
	var expenses, err1 = getUserInput("expenses: ")

	if err1 != nil {
		fmt.Println(err)
	}
	var taxRate, err2 = getUserInput("Taxes: ")

	if err2 != nil {
		fmt.Println(err)
	}

	//outputText(revenue, expenses, taxRate)
	var EBT, profit, ratio = calculate(revenue, expenses, taxRate)

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)
	storeResults(EBT, profit, ratio)

}

/*func outputText(revenue float64, expenses float64, taxRate float64) {

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")

	fmt.Scan(&expenses)

	fmt.Print("Enter the tax Rate: ")

	fmt.Scan(&taxRate)

}*/

func calculate(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {

	var EBT = revenue - expenses
	var profit = EBT - (1 - taxRate/100)
	var ratio = EBT / profit

	return EBT, profit, ratio

}

func storeResults(EBT, profit, ratio float64) {
	results := fmt.Sprint("EBT: %.1f\n Profit: %.1f\n Ratio: %.3f", EBT, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be positive")
	}
	return userInput, nil
}
