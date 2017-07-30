package main

import (
	"log"
	"os"

	_ "./search"
)

//execute before main
func init() {
	log.SetOutput(os.Stdout)
}

func main() {

}
