package main

import (
	"log"
	"os"

	_ "./search"
	"./matchers"
)

//execute before main
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
