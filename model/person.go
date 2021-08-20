package model

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func NewPerson(name string, age int, gender string) Person {
	return Person{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}
