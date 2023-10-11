package main // Define the package name

import (
	"bufio" // Import the bufio package for reading files
	"fmt"   // Import the fmt package for printing to the terminal
	"io"    // Import the io package for handling file I/O
	"os"    // Import the os package for accessing command-line arguments and files
)

// Define an interface for a file that can be opened and read
type file interface {
	openAndReadFile() (string, error)
}

// Define a type for a filename
type name string

// Define the main function
func main() {
	// Get the filename from the command-line arguments
	filename := os.Args[1]
	// Create a name object from the filename
	n := name(filename)
	// Call the printContent function with the name object as an argument
	printContent(n)
}

// Define the openAndReadFile method for the name type
func (n name) openAndReadFile() (string, error) {
	// Open the file with the name of the object
	file, err := os.Open(string(n))
	if err != nil {
		return "", err
	}
	// Defer closing the file until the end of the function
	defer file.Close()

	// Create a string variable to hold the contents of the file
	var content string
	// Create a bufio.Reader to read the file
	reader := bufio.NewReader(file)
	// Loop through the file line by line
	for {
		// Read the next line of the file
		line, err := reader.ReadString('\n')
		// If there is an error other than EOF, return the error
		if err != nil && err != io.EOF {
			return "", err
		}
		// Append the line to the content variable
		content += line
		// If we have reached the end of the file, break out of the loop
		if err == io.EOF {
			break
		}
	}

	// Return the contents of the file and any error that occurred
	return content, nil
}

// Define the printContent function
func printContent(f file) {
	// Call the openAndReadFile method on the file object
	content, err := f.openAndReadFile()
	// If there is an error, print the error message and return
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the contents of the file to the terminal
	fmt.Println(content)
}
