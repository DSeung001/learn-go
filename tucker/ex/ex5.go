package main

import "fmt"

type Animal interface {
	SetName(name string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d *Dog) GetName() string {
	return d.Name
}

func introduce(animal Animal) {
	fmt.Printf("안녕하세요, 저는 %s입니다.\n", animal.GetName())
}

func main() {
	dog := &Dog{}
	dog.SetName("멍멍이")
	introduce(dog)
}
