package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // Same like - contactInfo contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person{
		firstName: "Jim",
		lastName:  "Matthew",
		contactInfo: contactInfo{
			email:   "jim.matthew@example.com",
			zipCode: 90005, // even last property must have comma ','
		},
	}
	fmt.Println(alex)

	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Gotcha - 1

//func main(){
//	mySlice := []string{"Hi","There"}
//	updateSlice(mySlice)
//	fmt.Println(mySlice)
//}
//
//func updateSlice(s []string){
//	s[0] = "Bye"
//}
