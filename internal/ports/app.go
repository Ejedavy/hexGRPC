package ports

type APIPort interface {
	GetAddition(int32, int32) (int32, error)
	GetSubstraction(int32, int32) (int32, error)
	GetDivision(int32, int32) (int32, error)
	GetMultiplication(int32, int32) (int32, error)
}
