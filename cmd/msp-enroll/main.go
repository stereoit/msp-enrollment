package main

import (
	"fmt"
	"os"

	"github.com/stereoit.com/enrollment"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: msp-enroll <vin>")
		os.Exit(0)
	}

	vin := os.Args[1]
	wizard := enrollment.NewEnrollByVIN(vin)
	result := wizard.Enroll()

	if result.Code == enrollment.ResultCodeFail {
		fmt.Printf("Enrollment for VIN \"%v\" failed: %v\n", vin, result.ErrMsg)
		os.Exit(1)
	}

	fmt.Printf("Vehicle \"%v\" VIN: \"%v\" was enrolled successfully.\n", result.VehicleEntity.VehicleCode, vin)
}
