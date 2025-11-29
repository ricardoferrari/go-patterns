package main

import "fmt"

type Person struct {
	Name     string
	Address  string
	Postcode string
	City     string
	Position string
	Income   float64
}

type PersonBuilder struct {
	person *Person
}

type PersonJobBuilder struct {
	builder *PersonBuilder
}

type PersonAddressBuilder struct {
	builder *PersonBuilder
}

//Name method
func (pb *PersonBuilder) Called(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

//Address methods
func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{builder: pb}
}

func (ab *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	ab.builder.person.Address = address
	return ab
}
func (ab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	ab.builder.person.Postcode = postcode
	return ab
}
func (ab *PersonAddressBuilder) In(city string) *PersonBuilder {
	ab.builder.person.City = city
	return ab.builder
}

// Job methods
func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{builder: pb}
}

func (jb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	jb.builder.person.Position = position
	return jb
}

func (jb *PersonJobBuilder) Earning(income float64) *PersonBuilder {
	jb.builder.person.Income = income
	return jb.builder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func main() {
	person := NewPersonBuilder().
		Called("John Doe").
		Lives().
		At("123 Main St").
		WithPostcode("12345").
		In("Anytown").
		Works().
		AsA("Software Engineer").
		Earning(75000).
		Build()
	fmt.Println(person)
}
