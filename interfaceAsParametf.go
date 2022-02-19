package main

import (
	"strconv"
	"fmt"
)

type Employee interface {
	GetDetails() string
}

func GetUserDetails(emp Employee) {
	emp.GetDetails()
}

type Manager struct {
	Name        string
	Age         int
	Designation string
	Salary      int
}

type Lead struct {
	Name     string
	Age      int
	TeamSize string
	Salary   int
}

func (mgr Manager) GetDetails() string {
	return mgr.Name + " " + strconv.Itoa(mgr.Age)
}

func (ld Lead) GetDetails() string {
	return ld.Name + " " + strconv.Itoa(ld.Age)
}

func main() {
	newLead := Lead{Name: "Mayank", Age: 30, TeamSize: "30", Salary: 10}
	newManager := Manager{Name: "Mayank", Age: 30, Designation: "228", Salary: 10}
	var employeeInterface Employee

	employeeInterface = newManager
	fmt.Println((employeeInterface))
	fmt.Println((newManager))
	fmt.Println((newLead))
}
