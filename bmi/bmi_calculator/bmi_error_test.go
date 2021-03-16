package bmi_calculator

import (
	"testing"
)

func TestBmiMinusValueError(t *testing.T) {
	err := BmiError{BmiMinusValueError}
	if err.Error() != "値がマイナスです" {
		t.Error("\n実際： ", err, "\n理想： ", "値がマイナスです")
	}
}

func TestBmiTooBigValueError(t *testing.T) {
	err := BmiError{BmiTooBigValueError}
	if err.Error() != "値が大きすぎます" {
		t.Error("\n実際： ", err, "\n理想： ", "値が大きすぎます")
	}
}

func TestBmiTooSmallValueError(t *testing.T) {
	err := BmiError{BmiTooSmallValueError}
	if err.Error() != "値が小さすぎます" {
		t.Error("\n実際： ", err, "\n理想： ", "値が小さすぎます")
	}
}

func TestBmiUnexceptedValueError(t *testing.T) {
	err := BmiError{0}
	if err.Error() != "unknown error code" {
		t.Error("\n実際： ", err, "\n理想： ", "unknown error code")
	}
}
