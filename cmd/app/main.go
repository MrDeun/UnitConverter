package main

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println("Unit Converter - utility designed to convert units.\n Usage: unitconv (input_unit) (output_unit) (quantity)")
}

func main() {
	args := os.Args[1:]
	if args[0] == "-h" {
		Help()
		os.Exit(0)
	}
	if len(args) != 3 {
		fmt.Printf("Error: expected 3 arguments, received: %d\n", len(args))
		os.Exit(1)
	}

}
