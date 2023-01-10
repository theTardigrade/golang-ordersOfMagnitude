package ordersOfMagnitude

import (
	"math"
)

func Name(number int64) (value string, found bool) {
	if number < 0 {
		if number == math.MinInt64 {
			number = math.MaxInt64
		} else {
			number *= -1
		}
	}

	if number < 1e3 {
		if number < 10 {
			return "units", true
		}

		if number < 100 {
			return "tens", true
		}

		return "hundreds", true
	}

	if number < 1e6 {
		return nameHelper(number, 6, "thousands"), true
	}

	if number < 1e9 {
		return nameHelper(number, 9, "millions"), true
	}

	if number < 1e12 {
		return nameHelper(number, 12, "billions"), true
	}

	if number < 1e15 {
		return nameHelper(number, 15, "trillions"), true
	}

	if number < 1e18 {
		return nameHelper(number, 18, "quadrillions"), true
	}

	return
}

func nameHelper(number int64, upperBoundPowerOfTenExponent int64, baseName string) (value string) {
	if number < int64(math.Pow(10, float64(upperBoundPowerOfTenExponent-2))) {
		return baseName
	}

	if number < int64(math.Pow(10, float64(upperBoundPowerOfTenExponent-1))) {
		return "tens of " + baseName
	}

	return "hundreds of " + baseName
}
