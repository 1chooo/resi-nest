/*
Create Date: 2023/08/21
Author: @1chooo (Hugo ChunHo Lin)
Version: v0.0.1
*/

package main

import (
	"fmt"
	"resi_nest/resi/divisor"
	"resi_nest/resi/client"
)

func main() {
	fmt.Println("Hello Go")
	client.HelloClient()
	fmt.Println("Executing main.go in resi_nest")
	remainders := divisor.CalculateRemainders(100, 7)
	fmt.Println("Remainders:", remainders)
}
