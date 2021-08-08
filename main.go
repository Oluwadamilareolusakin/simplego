// package main

// import "github.com/oluwadamilareolusakin/calculator/operations"

// func main() {
//   operations.AdditionWithSteps(1, 2, 3, 4, 5)
// }

package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupMX("fixus.com")

		fmt.Println(*iprecords[0])
	}
