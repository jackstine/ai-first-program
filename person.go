package main

import "fmt"

// Person represents a person with basic information
type Person struct {
	Name  string
	Age   int
	Email string
}

// NewPerson creates a new Person instance
func NewPerson(name string, age int, email string) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

// String returns a string representation of the Person
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}

// DemonstratePerson creates and displays an example person
func DemonstratePerson() {
	person := NewPerson("John Doe", 30, "john.doe@example.com")
	fmt.Println("Person example:")
	fmt.Println(person)
}