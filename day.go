package childcare

import "time"

type Day struct {
	Id  int
	Date time.Time
	Child  Child
    Room Room
    Facility Facility
    Caregiver Caregiver
    Events []Event
    Notes string
}

