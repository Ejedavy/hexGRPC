package ports

type ArithmeticPort interface {
	Addition(int32, int32) (int32, error)
	Substraction(int32, int32) (int32, error)
	Multiplication(int32, int32) (int32, error)
	Division(int32, int32) (int32, error)
}
