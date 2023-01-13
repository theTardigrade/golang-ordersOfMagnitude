package ordersOfMagnitude

import (
	"math"
)

var (
	nameValues = map[int64]string{
		1:  "units",
		2:  "tens",
		3:  "hundreds",
		6:  "thousands",
		9:  "millions",
		12: "billions",
		15: "trillions",
		18: "quadrillions",
	}
)

func Name(number int64) (value string, found bool) {
	if number < 0 {
		if number == math.MinInt64 {
			number = math.MaxInt64
		} else {
			number *= -1
		}
	}

	for i := int64(1); i <= 3; i++ {
		if value, found = nameHelper(number, i); found {
			return
		}
	}

	for i := int64(6); i <= 18; i += 3 {
		if value, found = nameHelper(number, i); found {
			return
		}
	}

	return
}

func nameHelper(number int64, upperBoundPowerOfTenExponent int64) (value string, found bool) {
	upperBound := int64(math.Pow(10, float64(upperBoundPowerOfTenExponent)))

	if number >= upperBound {
		return
	}

	if upperBoundPowerOfTenExponent <= 3 || number < int64(math.Pow(10, float64(upperBoundPowerOfTenExponent-2))) {
		value = nameValues[upperBoundPowerOfTenExponent]
	} else if number < int64(math.Pow(10, float64(upperBoundPowerOfTenExponent-1))) {
		value = nameValues[2] + " of " + nameValues[upperBoundPowerOfTenExponent]
	} else {
		value = nameValues[3] + " of " + nameValues[upperBoundPowerOfTenExponent]
	}

	found = true

	return
}
