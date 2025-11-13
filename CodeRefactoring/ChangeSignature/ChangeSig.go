package main

import "fmt"

func computeSalary(base float64, bonus float64, tax float64) float64 {
	return base + bonus - tax
}

func displayEmployeeSalary(name string, monthlySalary float64) {
	fmt.Printf("%s earns %.2f per month.\n", name, monthlySalary)
}

func main() {
	salary := computeSalary(30000, 5000, 2000)
	displayEmployeeSalary("Erick", salary)
}
