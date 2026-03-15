package main

import "github.com/solace06/AnraAssignment/api"

func main() {
	//router initialized
	router := api.NewRouter()

	//server initialized
	router.Run(":8080")
}
