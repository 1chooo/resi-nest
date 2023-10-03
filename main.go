/*
Create Date: 2023/08/21
Author: @1chooo (Hugo ChunHo Lin)
Version: v0.0.1
*/

package main

import (
	"fmt"

	"github.com/1chooo/resi_nest/resi/client"
	"github.com/1chooo/resi_nest/resi/server"
)

func main() {
	fmt.Println("Hello Go")
	client.HelloClient()
	fmt.Println("Executing main.go in resi_nest")
	server.HelloServer()
}
