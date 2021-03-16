package bmi_calculator

import (
	"errors"
)

func GetBmiFromHeightAndWeight(height float64, weight float64) (float64, error) {
	if weight < 20 {
		return 0, errors.New("体重がかるすぎます")
	}
	if height < 100 {
		return 0, errors.New("身長が低すぎます")
	}
	var res float64 = weight / height / height
	return res, nil
}
