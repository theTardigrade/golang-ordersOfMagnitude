# golang-ordersOfMagnitude

This package exposes a single function, ```Name```, which describes the order of magnitude of a number.

[![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-ordersOfMagnitude)](https://goreportcard.com/report/github.com/theTardigrade/golang-ordersOfMagnitude)

## Example

```golang
package main

import (
	"fmt"

	ordersOfMagnitude "github.com/theTardigrade/golang-ordersOfMagnitude"
)

func main() {
	const exampleNumber = 12_345_678_912

	fmt.Printf(
		"The number %d has an order of magnitude in the %s.\n",
		exampleNumber,
		ordersOfMagnitude.Name(exampleNumber),
	)
}
```

Running the example code above should produce the following output:

```
The number 12345678912 has an order of magnitude in the tens of billions.
```