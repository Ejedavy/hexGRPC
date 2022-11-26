package ports

type DBPort interface {
	CloseConnection()
	AddHistory(answer int32, operation string) error
}
