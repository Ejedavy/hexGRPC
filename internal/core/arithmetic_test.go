package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := New()
	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("Error!. Required %v found %v", nil, err)
	}
	require.Equal(t, answer, int32(2))
}

func TestSubstraction(t *testing.T) {
	arith := New()
	answer, err := arith.Subtraction(2, 1)
	if err != nil {
		t.Fatalf("Error!. Required %v found %v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}

func TestMultiplication(t *testing.T) {
	arith := New()
	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("Error!. Required %v found %v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}

func TestDivision(t *testing.T) {
	arith := New()
	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("Error!. Required %v found %v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}
