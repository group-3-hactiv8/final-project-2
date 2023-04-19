package main

import (
	"final-project-2/handlers"
)

const port = ":8080"

func main() {
	r := handlers.StartApp()

	r.Run(port)
}
