package main

import "fmt"

func main() {
	var emp Employee = &ContractualEmployee{
		EmployeeDetails: EmployeeDetails{
			EmployeeID: 123,
			Name:       "abc",
		},
		NumberOfWorkingHrs: 99,
	}

	emp.PrintDetails()
}

type Employee interface {
	PrintDetails()
	CalculateSalary()
}

type ContractualEmployee struct {
	EmployeeDetails
	NumberOfWorkingHrs int
}

type SalariedEmployee struct {
	EmployeeDetails
	YearsOfExperience int
}

type EmployeeDetails struct {
	EmployeeID int
	Name       string
}

func (c *ContractualEmployee) PrintDetails() {
	fmt.Println(c.Name)
	fmt.Println(c.EmployeeID)
	fmt.Println(c.NumberOfWorkingHrs)
}

func (c *ContractualEmployee) CalculateSalary() {

}
