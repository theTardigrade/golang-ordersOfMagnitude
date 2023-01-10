# golang-ordersOfMagnitude

## Example

```golang
package main

import (
	"fmt"

	ordersOfMagnitude "github.com/theTardigrade/golang-ordersOfMagnitude"
)

func main() {
	const exampleNumber = 12_345_678_912

	if name, found := ordersOfMagnitude.Name(exampleNumber); found {
		fmt.Printf(
			"The number %d has an order of magnitude in the %s.\n",
			exampleNumber,
			name,
		)
	}
}
```