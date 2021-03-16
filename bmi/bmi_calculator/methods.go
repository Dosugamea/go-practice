package bmi_calculator

func GetBmiFromHeightAndWeight(height Height, weight Weight) (Bmi, error) {
	if weight < 0 || height < 0 {
		return 0, BmiError{BmiMinusValueError}
	}
	if weight <= 20 {
		return 0, BmiError{BmiTooSmallValueError}
	}
	if weight >= 300 {
		return 0, BmiError{BmiTooBigValueError}
	}
	if height <= 100 {
		return 0, BmiError{BmiTooSmallValueError}
	}
	if height >= 300 {
		return 0, BmiError{BmiTooBigValueError}
	}
	var res Bmi = Bmi(
		float64(weight) / (float64(height)/100 + float64(height)/100),
	)
	return res, nil
}
