package main

import (
	"fmt"
	"math"
	"notificator/main/lib"
	"os/exec"
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

	matrixCommanderPath := "/root/matrix-commander/venv/bin/matrix-commander"
	msg := fmt.Sprintf(`Health Factor: %v`, healthFactor)

	cmd := exec.Command(matrixCommanderPath, "-m", msg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error with sending msg: %s\n", err)
	}

	fmt.Println(string(output))
}
