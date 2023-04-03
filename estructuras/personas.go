package estructuras

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Id int
	Person
	Position string
}

func (e *Employee) PrintEmployee() {
	fmt.Println(e)
}
