package main
import "fmt"
type person struct {
	firstname string
	lastname string
	//contactInfo contactInfo = contactInfo
	contactInfo
}
type contactInfo struct{
	email string
	zipcode int
}

func main (){

	// alex := person{"Alex", "Anderson"}
	// // vedant := person{firstname: "Alex",lastname: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstname = "Alex"
	// alex.lastname = "Anderson"
	// alex.contact.email = "alex.anderson@go.com"
	// alex.contact.zipcode = 32608

	jim := person{
		firstname:"Jim",
		lastname:"Party",
		contactInfo : contactInfo{
			email:"jima@go.com",
			zipcode:60886,
		},
	}
	// jimPointer := &jim
	// fmt.Println(jimPointer)
	jim.updateName("Jimmy")
	jim.print()

	// Arrays are point by reference
}

func (p person) print(){
	fmt.Println(p)
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstname = newFirstName
}