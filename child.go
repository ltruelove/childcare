package childcare

import "fmt"

type Child struct {
	Id                  int
	FirstName, LastName string
}

func (c Child) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}
