package load

import (
	"bufio"
	"fmt"
	"os"
)

func Path(l *Load) {
	file, err := os.Open(l.Path)
	if err != nil {
		l.Errors["-p"].GetErrors(1)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and print each line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		l.Errors["-p"].GetErrors(2)
	}
}
