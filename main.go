package main

import "todolist/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
