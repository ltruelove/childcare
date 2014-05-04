package childcare

import "time"

type Event struct {
	Id  int
    Type string
	Date time.Time
	Child  Child
    Caregiver Caregiver
    Events []Event
}

