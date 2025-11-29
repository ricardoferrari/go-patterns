package main

import "fmt"

type Person struct {
	Name, Address, City string
}

type PersonModifier func(*Person)

type PersonBuilder struct {
	actions []PersonModifier
}

func (pb *PersonBuilder) Called(name string) *PersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		fmt.Println("Adding name...")
		p.Name = name
	})
	return pb
}

func (pb *PersonBuilder) LivesAt(address, city string) *PersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		fmt.Println("Adding address and city...")
		p.Address = address
		p.City = city
	})
	return pb
}

func (pb *PersonBuilder) Build() *Person {
	p := &Person{}
	for _, action := range pb.actions {
		action(p)
	}
	return p
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func main() {
	// Example usage:
	personBuilder := NewPersonBuilder().
		Called("John Doe").
		LivesAt("123 Main St", "Anytown")
	// This lazily initializes the Person object but does not build it until Build() is called.

	fmt.Println("Building person lazily because of deferred execution of actions...")
	person := personBuilder.Build()
	// person is now a *Person with Name "John Doe", Address "123 Main St", City "Anytown"
	fmt.Println("Person built:")
	fmt.Println(*person)

}
