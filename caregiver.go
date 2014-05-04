package childcare

import "fmt"

type Caregiver struct {
	Id                  int
	FirstName, LastName string
	Children            []Child
}

func (c Caregiver) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}
