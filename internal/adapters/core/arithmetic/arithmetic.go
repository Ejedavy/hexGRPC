package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (this Adapter) Addition(a int32, b int32) (int32, error) {
	c := a + b
	return c, nil
}

func (this Adapter) Substraction(a int32, b int32) (int32, error) {
	c := a - b
	return c, nil
}

func (this Adapter) Multiplication(a int32, b int32) (int32, error) {
	c := a * b
	return c, nil
}

func (this Adapter) Division(a int32, b int32) (int32, error) {
	c := (a / b)
	return c, nil
}
