package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 || span > len(digits) {
		return 0, errors.New("Span should be greatger than 0!")
	}
	if span == 0 {
		return 1, nil
	}
	var mPrdct int64 = 0
	for i := 0; i <= len(digits)-span; i++ {
		cPrdct := int64(1)
		for j := 0; j < span; j++ {
			digit, err := strconv.Atoi(string(digits[i+j]))
			if err != nil {
				return 0, errors.New("Incorrect input! Check it!")
			}
			cPrdct *= int64(digit)
		}
		if cPrdct > mPrdct {
			mPrdct = cPrdct
		}
	}
	return mPrdct, nil
}
