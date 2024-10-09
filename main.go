package main

import (
	"fmt"
	"math"
	"notificator/main/lib"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	matrixCommanderPath := os.Getenv("MATRIX_COMMANDER_PATH")
	if matrixCommanderPath == "" {
		fmt.Println("Matrix commander path not found in your environment vars.")
	}
	credentialsFilePath := os.Getenv("MATRIX_CREDENTIALS_FILE_PATH")
	if credentialsFilePath == "" {
		fmt.Println("Matrix credentials file path not found in your environment vars.")
	}
	storePath := os.Getenv("MATRIX_STORE_PATH")
	if storePath == "" {
		fmt.Println("Matrix store directory path not found in your environment vars.")
	}
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

	msg := fmt.Sprintf(`Health Factor: %v`, healthFactor)

	cmd := exec.Command(matrixCommanderPath, "-m", msg, "-c", credentialsFilePath, "-s", storePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error with sending msg: %s\n", err)
	}

	fmt.Println(string(output))
}
