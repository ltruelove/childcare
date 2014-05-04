EXAMPLE USAGE:

package main

import (
	"fmt"
    "time"
	//"github.com/ltruelove/newmath"
	"github.com/ltruelove/childcare"
)

func main() {
	//children belong as a group to many things
	child := childcare.Child{0, "Some", "Kid"}
	child2 := childcare.Child{1, "Other", "Child"}
	child3 := childcare.Child{2, "Third", "One"}
	var childSlice = make([]childcare.Child, 2)

	childSlice[0] = child
	childSlice[1] = child2

	//caregiver/teacher/whatever you want to call them
	caregiver := childcare.Caregiver{1, "Carol", "Giver", childSlice}

	caregiver.Children = append(caregiver.Children, child3)

	fmt.Printf("\nCaregiver: %s\n", caregiver.FullName())
	fmt.Printf("%s children:\n", caregiver.FullName())

	for i := 0; i < len(caregiver.Children); i++ {
		giverChild := caregiver.Children[i]

		fmt.Printf("%s\n", giverChild.FullName())
	}

	//rooms are a part of a facility
	room1 := childcare.Room{0, "Infants", caregiver, caregiver.Children}
	rooms := make([]childcare.Room, 1)
	rooms[0] = room1

	//slice of caregivers that work at the facility
	caregivers := make([]childcare.Caregiver, 1)
	caregivers[0] = caregiver

	//facility is the center or home
	facility := childcare.Facility{0,
		"Day Care Place",
		"123 test street",
		"",
		"",
		"",
		"",
		"",
		caregiver.Children, //in real life we'd use a more comprehensive list of children
		caregivers,
		rooms}

	fmt.Printf("Day Care facility %s has %d children, %d caregiver(s), and %d rooms\n",
		facility.Name,
		len(facility.Children),
		len(facility.Caregivers),
		len(facility.Rooms))

	//events happen on days
	diaperEvent := childcare.Event{0,
		"Diaper change, wet",
		time.Now(),
		caregiver.Children[0],
		caregiver,
		"Note: applied cream"}

	eventSlice := make([]childcare.Event, 1)
	eventSlice[0] = diaperEvent

	//a day represents a child's typical day at the facility
	childDay := childcare.Day{0,
		time.Now(),
		caregiver.Children[0],
		facility.Rooms[0],
		facility,
		caregiver,
		eventSlice,
		"Note: Some Kid had a great day!"}

	fmt.Printf("%s %s had %d events today\n",
		childDay.Date.Format("Jan 2, 2006"),
		childDay.Child.FullName(),
		len(childDay.Events))
}

