/*
Package app represents files that sum up to the application layer of the software.
This layer **strongly** depends on the business layer but weakly depend on created Ports
that assist the application in fulfilling its functionalities.
*/
package app

import (
	"github.com/ejedavy/hexGRPC/internal/core"
	"github.com/ejedavy/hexGRPC/internal/ports"
)

type ArithmeticApp struct {
	arith core.Arithmetic
	db    ports.DBPort
}

func New(db ports.DBPort) *ArithmeticApp {
	return &ArithmeticApp{arith: core.Arithmetic{}, db: db}
}

func (api ArithmeticApp) GetAddition(a int32, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddHistory(answer, "Addition")
	return answer, err
}

func (api ArithmeticApp) GetDivision(a int32, b int32) (int32, error) {
	answer, err := api.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddHistory(answer, "Division")
	return answer, err
}

func (api ArithmeticApp) GetMultiplication(a int32, b int32) (int32, error) {
	answer, err := api.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddHistory(answer, "Multiplication")
	return answer, err
}

func (api ArithmeticApp) GetSubtraction(a int32, b int32) (int32, error) {
	answer, err := api.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddHistory(answer, "GetSubtraction")
	return answer, err
}
