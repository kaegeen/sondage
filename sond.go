package main

import (
	"fmt"
)

func main() {

	var connectMultipleTimes float64 = 10.2
	var noConnectionPrinciple float64 = 21.4
	var noConnectionNotNecessary float64 = 45.1
	var connectSometimes float64 = 23.3

	var connectDueToObligation float64 = 31.0
	var connectDueToWorkloadFear float64 = 26.0

	fmt.Println("French Workers: Connection During Vacations")
	fmt.Printf("10.2%% connect multiple times per day\n")
	fmt.Printf("21.4%% do not connect at all (out of principle)\n")
	fmt.Printf("45.1%% do not connect (because it wasn't necessary)\n")
	fmt.Printf("23.3%% connect from time to time\n")

	var response string
	fmt.Print("\nWould you like to see why some connect during vacations? (yes/no): ")
	fmt.Scanln(&response)

	if response == "yes" {

		fmt.Println("\nReasons for connecting during vacations:")
		fmt.Printf("31%% connect due to professional obligations\n")
		fmt.Printf("26%% connect due to fear of workload upon return\n")
	} else {
		fmt.Println("\nThank you! Have a great day!")
	}
}
