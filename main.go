package main

import (
	"fmt"
	"math"
	"notificator/main/lib"
	"os"
	"os/exec"
	"time"

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

	for true {
		position, err := lib.QueryPosition()
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
		}

		totalDeposit, totalBorrow, err := lib.ProcessSum(position)
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
		}

		fmt.Printf(`Total deposits: %v\n`, totalDeposit)
		fmt.Printf(`Total borrows: %v\n`, totalBorrow)

		const LIQUIDATION_TRESHOLD_WBTC = 0.78
		healthFactor := math.Round(float64((totalDeposit*LIQUIDATION_TRESHOLD_WBTC)/totalBorrow)*100) / 100

		println("Health factor: ", healthFactor)

		timeZone := time.FixedZone("GMT+2", 2*60*60)
		currentTime := time.Now().In(timeZone)
		msg := fmt.Sprintf(`Health Factor: %v @ %s`, healthFactor, currentTime.Format("2006-01-02 15:04:05"))

		cmd := exec.Command(matrixCommanderPath, "-m", msg, "-c", credentialsFilePath, "-s", storePath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error with sending msg: %s\n", err)
		}

		fmt.Println(string(output))

		time.Sleep(60 * time.Minute)
	}
}
