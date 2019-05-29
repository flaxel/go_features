package employee

import (
	"fmt"
)

// employee is not exported because the name is written in lower cases
type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// export this function because the name is written in a capital letter
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
