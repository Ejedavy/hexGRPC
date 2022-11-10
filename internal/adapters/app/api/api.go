package api

import (
	"hex/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DBPort
}

func NewAdapter(a ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{arith: a, db: db}
}

func (api Adapter) GetAddition(a int32, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	api.db.AddHistory(answer, "Addition")
	return answer, err
}

func (api Adapter) GetDivision(a int32, b int32) (int32, error) {
	answer, err := api.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	api.db.AddHistory(answer, "Division")
	return answer, err
}

func (api Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	answer, err := api.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	api.db.AddHistory(answer, "Multiplication")
	return answer, err
}

func (api Adapter) GetSubstraction(a int32, b int32) (int32, error) {
	answer, err := api.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}
	api.db.AddHistory(answer, "Substraction")
	return answer, err
}
