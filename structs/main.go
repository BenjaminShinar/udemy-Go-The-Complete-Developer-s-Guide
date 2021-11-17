package main

import "fmt"

type person_1 struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateNameFails(newFirstName string) {
	p.firstName = newFirstName
	//p.print() // name is changed here, but not outside!
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	alex := person_1{"Alex", "Anderson"}                     //positional argument
	jack := person_1{firstName: "jack", lastName: "jackson"} //names arguments
	var jon person_1                                         // zero value initialization
	fmt.Println(alex, jack, jon)

	fmt.Printf("%v , %+v \n", alex, alex) // field and value names
	jon.firstName = "jon"
	// jon.lastName := "jonson" //nope, can't use :=
	// jon["lastName"] = "jonson" //nope, this isn't a javascript object!
	fmt.Println(jon)

	jim := person{
		firstName: "jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "party@jim.now",
			zipCode: 94000,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%+v\n", jim)
	jim.updateNameFails("jimmy")
	jim.print()
	pointerToJim := &jim
	pointerToJim.updateName("jimmy2") //call from pointer
	jim.print()
	jim.updateName("jimmy3") //call from value, also works
	jim.print()

}
