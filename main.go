package main // Define the package name

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file specified in the command-line arguments
	f, err := os.Open(os.Args[1])
	if err != nil {
		// If there was an error opening the file, print the error message and exit
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Copy the contents of the file to the standard output
	io.Copy(os.Stdout, f)
}
