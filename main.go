package main

import (
	"fmt"
	"math"
	"notificator/main/lib"
)

func main() {
	position, err := lib.QueryPosition()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	fmt.Print(position + "\n")

	totalDeposit, totalBorrow, err := lib.ProcessSum(position)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	fmt.Printf(`Total deposits: %v\n`, totalDeposit)
	fmt.Printf(`Total borrows: %v\n`, totalBorrow)

	healthFactor := math.Round(float64((totalDeposit*0.78)/totalBorrow)*100) / 100

	println("Health factor: ", healthFactor)
}
