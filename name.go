package ordersOfMagnitude

import (
	"fmt"
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

func init() {
	for i := int64(18); i > 3; i -= 3 {
		currentNameValue := nameValues[i]

		nameValues[i-2] = currentNameValue
		nameValues[i-1] = nameValues[2] + " of " + currentNameValue
		nameValues[i] = nameValues[3] + " of " + currentNameValue
	}

	fmt.Println(nameValues)
}

func Name(number int64) (value string, found bool) {
	if number < 0 {
		if number == math.MinInt64 {
			number = math.MaxInt64
		} else {
			number *= -1
		}
	}

	for i := int64(1); i <= 18; i++ {
		upperBound := int64(math.Pow(10, float64(i)))

		if number < upperBound {
			value, found = nameValues[i], true
			return
		}
	}

	return
}
