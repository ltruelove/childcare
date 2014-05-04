package childcare

//import "fmt"

type Facility struct {
	Id                                             int
	Name, Street, Street2, City, State, Zip, Phone string
	Children                                       []Child
	Caregivers                                     []Caregiver
	Rooms                                          []Room
}
