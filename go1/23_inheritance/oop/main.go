package main

import (
	"23_inheritance/oop/employee"
)

// go is not pure oo
// go provides only classes
func main() {
	// CREATE EMPLOYEE
	/*e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()*/

	// CREATE EMPTY EMPLOYEE - go does not support constructors
	/*var e2 employee.Employee
	e2.LeavesRemaining()*/

	// CREATE EMPLOYEE WITH 'NEW' METHOD
	e3 := employee.New("John", "Doe", 30, 20)
	e3.LeavesRemaining()
}
