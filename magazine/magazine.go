//Package magazine define a structure
package magazine

//Address is a address structure
type Address struct {
	Street     string
	PostalCode int
	City       string
	State      string
}

//Subscriber is a subscription structure
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Employee is a employee structure
type Employee struct {
	Name   string
	Salary float64
	Address
}
