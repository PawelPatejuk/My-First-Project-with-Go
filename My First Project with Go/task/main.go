package main

import "fmt"

func main() {
	// Write your code here
	fmt.Println(`Earned amount:
Bubblegum: $202
Toffee: $118
Ice cream: $2250
Milk chocolate: $1680
Doughnut: $1075
Pancake: $80
\nIncome: $5405`)

	var staff, other int
	fmt.Println("Staff expenses:")
	fmt.Scan(&staff)
	fmt.Println("Other expenses:")
	fmt.Scan(&other)
	fmt.Println("Net income: $%d", 5405-staff-other)
}
