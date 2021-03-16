package bmi_calculator

type BmiErrorCode uint

type BmiError struct {
	Code BmiErrorCode
}

const (
	BmiMinusValueError    BmiErrorCode = 1
	BmiTooBigValueError   BmiErrorCode = 2
	BmiTooSmallValueError BmiErrorCode = 3
)

func (e BmiError) Error() string {
	switch e.Code {
	case BmiMinusValueError:
		return "値がマイナスです"
	case BmiTooBigValueError:
		return "値が大きすぎます"
	case BmiTooSmallValueError:
		return "値が小さすぎます"
	default:
		return "unknown error code"
	}
}
