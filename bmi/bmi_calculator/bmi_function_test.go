package bmi_calculator

import (
	"testing"
)

func TestBmiWithMinusHeightAndProperWeight(t *testing.T) {
	var height Height = -160
	var weight Weight = 47
	_, err := GetBmiFromHeightAndWeight(height, weight)
	if err == nil {
		t.Error("\n実際： ", nil, "\n理想： ", BmiMinusValueError)
		return
	}
	bmi_err, ok := err.(BmiError)
	if !ok {
		t.Error("\n実際： Error", "\n理想： ", BmiMinusValueError)
		return
	}
	if bmi_err.Code != BmiMinusValueError {
		t.Error("\n実際： ", bmi_err.Code, "\n理想： ", BmiMinusValueError)
		return
	}
}

func TestBmiWithProperHeightAndMinusWeight(t *testing.T) {
	var height Height = 160
	var weight Weight = -47
	_, err := GetBmiFromHeightAndWeight(height, weight)
	if err == nil {
		t.Error("\n実際： ", nil, "\n理想： ", BmiMinusValueError)
		return
	}
	bmi_err, ok := err.(BmiError)
	if !ok {
		t.Error("\n実際： Error", "\n理想： ", BmiMinusValueError)
		return
	}
	if bmi_err.Code != BmiMinusValueError {
		t.Error("\n実際： ", bmi_err.Code, "\n理想： ", BmiMinusValueError)
		return
	}
}

func TestBmiWithProperHeightAndProperWeight(t *testing.T) {
	var height Height = 160
	var weight Weight = 47
	res, err := GetBmiFromHeightAndWeight(height, weight)
	if err != nil {
		t.Error("\n実際： ", err, "\n理想： ", nil)
		return
	}
	if res == 0 {
		t.Error("\n実際： ", res, "\n理想： ", nil)
		return
	}
}

func TestBmiWithBigHeightAndProperWeight(t *testing.T) {
	var height Height = 500
	var weight Weight = 47
	_, err := GetBmiFromHeightAndWeight(height, weight)
	if err == nil {
		t.Error("\n実際： ", err, "\n理想： ", nil)
		return
	}
	bmi_err, ok := err.(BmiError)
	if !ok {
		t.Error("\n実際： Error", "\n理想： ", BmiTooBigValueError)
		return
	}
	if bmi_err.Code != BmiTooBigValueError {
		t.Error("\n実際： ", bmi_err.Code, "\n理想： ", BmiTooBigValueError)
		return
	}
}

func TestBmiWithProperHeightAndBigWeight(t *testing.T) {
	var height Height = 160
	var weight Weight = 470
	_, err := GetBmiFromHeightAndWeight(height, weight)
	if err == nil {
		t.Error("\n実際： ", err, "\n理想： ", nil)
		return
	}
	bmi_err, ok := err.(BmiError)
	if !ok {
		t.Error("\n実際： Error", "\n理想： ", BmiTooBigValueError)
		return
	}
	if bmi_err.Code != BmiTooBigValueError {
		t.Error("\n実際： ", bmi_err.Code, "\n理想： ", BmiTooBigValueError)
		return
	}
}

func TestBmiWithSmallHeightAndProperWeight(t *testing.T) {
	var height Height = 5
	var weight Weight = 47
	_, err := GetBmiFromHeightAndWeight(height, weight)
	if err == nil {
		t.Error("\n実際： ", err, "\n理想： ", nil)
		return
	}
	bmi_err, ok := err.(BmiError)
	if !ok {
		t.Error("\n実際： Error", "\n理想： ", BmiTooSmallValueError)
		return
	}
	if bmi_err.Code != BmiTooSmallValueError {
		t.Error("\n実際： ", bmi_err.Code, "\n理想： ", BmiTooSmallValueError)
		return
	}
}

func TestBmiWithProperHeightAndSmallWeight(t *testing.T) {
	var height Height = 160
	var weight Weight = 4
	_, err := GetBmiFromHeightAndWeight(height, weight)
	if err == nil {
		t.Error("\n実際： ", err, "\n理想： ", nil)
		return
	}
	bmi_err, ok := err.(BmiError)
	if !ok {
		t.Error("\n実際： Error", "\n理想： ", BmiTooSmallValueError)
		return
	}
	if bmi_err.Code != BmiTooSmallValueError {
		t.Error("\n実際： ", bmi_err.Code, "\n理想： ", BmiTooSmallValueError)
		return
	}
}
