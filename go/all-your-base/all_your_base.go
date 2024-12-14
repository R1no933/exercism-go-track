package allyourbase

import "fmt"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	if len(inputDigits) == 0 {
		inputDigits = append(inputDigits, 0)
	}
	var res []int
	var decNum int
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		decNum = decNum*inputBase + digit
	}
	if decNum == 0 {
		res = append(res, 0)
	}
	for decNum > 0 {
		d := decNum % outputBase
		res = append(res, 0)
		copy(res[1:], res[0:])
		res[0] = d
		decNum = decNum / outputBase
	}
	return res, nil
}
