package main

import (
	"fmt"
	"strconv"
)

// Создание простого интерфейса Employee
type Employee interface {
	GetDetails() string
	GetEmployeeSalary() int
}

// Создание нового типа Manager, который содержит все функции, требующиеся интерфейсу Employee
type Manager struct {
	Name        string
	Age         int
	Designation string
	Salary      int
}


// Создание нового типа Lead, который содержит все функции, требующиеся интерфейсу Employee
type Lead struct {
	Name string
	Age int
	TeamSize string 
	Salary int
  }

func (ld Lead) GetDetails() string {
	return ld.Name + " " + strconv.Itoa(ld.Age)
}

func (ld Lead) GetEmployeeSalary() int {
	return ld.Salary
  }

func (mgr Manager) GetDetails() string {
	return mgr.Name + " " + strconv.Itoa(mgr.Age)
}
func (mgr Manager) GetEmployeeSalary() int {
	return mgr.Salary
} // Создание нового объекта типа Manager

func main() {
	newManager := Manager{Name: "Mayank",
		Age: 30, Designation: "Developer", Salary: 10}
	newLead := Lead{Name: "Mayank", 
		Age: 30, TeamSize: "30", Salary: 10}
	var employeeInterface Employee // Создана новая переменная типа Employee
	employeeInterface = newManager // Объект Manager присвоен интерфейсному типу, потому что контракт интерфейса выполнен
	employeeInterface.GetDetails() // Вызов функций, принадлежащих интерфейсу Employee
	fmt.Println(employeeInterface.GetDetails())
	employeeInterface = newLead
	fmt.Println(employeeInterface.GetDetails())
}
