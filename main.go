package main

import (
	"fmt"

	"github.com/treinalinux/api-go-rest/routes"
)

// main start the server
func main() {
	fmt.Println("Starting the server with Go...")
	routes.HandleRequest()
}
