package childcare

//import "fmt"

type Room struct {
	Id       int
	Name     string
	Lead     Caregiver
	Children []*Child
}
