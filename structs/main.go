package main

import "fmt"

type contactInfo struct{
	emailId string
	phoneNo int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// alex.firstName="Neetu"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName :"Jim",
		lastName :"Party",
		contact : contactInfo{
			emailId :"jim3232@gmail.com",
			phoneNo : 3787382928,
		},
	}

	// jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newName string){
	(*pointerToPerson).firstName = newName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}