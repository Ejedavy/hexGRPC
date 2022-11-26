/*
Package core contains classes / entities that sum up to the business layer
of the application.
*/
package core

// Arithmetic is the business logic of the software.
//
// It is responsible for performing the needed operation without any dependency on
// outer libraries or framework
type Arithmetic struct {
}

func New() *Arithmetic {
	return &Arithmetic{}
}

func (Arithmetic) Addition(a int32, b int32) (int32, error) {
	c := a + b
	return c, nil
}

func (Arithmetic) Subtraction(a int32, b int32) (int32, error) {
	c := a - b
	return c, nil
}

func (Arithmetic) Multiplication(a int32, b int32) (int32, error) {
	c := a * b
	return c, nil
}

func (Arithmetic) Division(a int32, b int32) (int32, error) {
	c := a / b
	return c, nil
}
