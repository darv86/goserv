package test

import "log"

type Person struct {
	Name string
	Age  int
}

func (person *Person) Hi() {
	log.Printf("'Hi': %s says", person.Name)
}

func (person *Person) AgeIs() {
	log.Printf("%s's age is %v", person.Name, person.Age)
}

func (person *Person) SetAge(age int) {
	person.Age = age
	log.Printf("%s's age has changed to %v", person.Name, person.Age)
}

type User struct {
	Person
	Status string
}
