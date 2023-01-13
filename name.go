package ordersOfMagnitude

import (
	"math"
)

var (
	nameValues = map[int64]string{
		1:  "units",
		2:  "tens",
		3:  "hundreds",
		4:  "thousands",
		7:  "m",
		10: "b",
		13: "tr",
		16: "quadr",
		19: "quint",
	}
)

func init() {
	for i := int64(4); i <= 19; i += 3 {
		currentNameValue := nameValues[i]

		if i >= 7 {
			currentNameValue += "illions"
			nameValues[i] = currentNameValue
		}

		if i <= 16 {
			for j := int64(1); j <= 2; j++ {
				nameValues[i+j] = nameValues[j+1] + " of " + currentNameValue
			}
		}
	}
}

func Name(number int64) string {
	if number < 0 {
		if number == math.MinInt64 {
			number = math.MaxInt64
		} else {
			number *= -1
		}
	}

	for i := int64(1); i <= 18; i++ {
		if upperBound := int64(math.Pow(10, float64(i))); number < upperBound {
			return nameValues[i]
		}
	}

	return nameValues[19]
}
