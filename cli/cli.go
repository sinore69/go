package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define command-line flags
	name := flag.String("name", "World", "Specify your name")
	greet := flag.Bool("greet", false, "Whether to greet the user")

	// Parse command-line flags
	flag.Parse()

	// Check if the greet flag is provided
	if *greet {
		fmt.Printf("Hello, %s!\n", *name)
	} else {
		fmt.Printf("You did not choose to greet %s.\n", *name)
	}
}
