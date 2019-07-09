package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	mealCost, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	tips, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	tax, _ := strconv.ParseFloat(scanner.Text(), 64)

	var totalPrice = mealCost + mealCost*tips/100 + mealCost*tax/100

	fmt.Printf("%.0f", totalPrice)

}
