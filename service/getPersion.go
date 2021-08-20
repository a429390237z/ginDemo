package service

import (
	"fmt"
	"ginDemo/model"
)

type People struct {
	Persons []model.Person `json:"persons"`
	Count   int            `json:"count"`
}

func NewPeople() *People {
	return &People{
		Persons: []model.Person{},
		Count:   0,
	}
}

func (p *People) AddPerson(t model.Person) {
	t.Id = p.Count + 1
	p.Persons = append(p.Persons, t)
	p.Count += 1
}

func (p *People) DelPerson(id int) {
	n := -1
	for i, v := range p.Persons {
		if v.Id == id {
			n = i
			break
		}
	}
	fmt.Println(n)
}

func (p *People) ShowCount() int {
	return p.Count
}
