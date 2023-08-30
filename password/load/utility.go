package load

import (
	"bufio"
	"os"
	"strings"

	"github.com/cetinboran/gosec/password/create"
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
		SaveInputs(l, CheckLine(l, line))
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		l.Errors["-p"].GetErrors(2)
	}
}

func CheckLine(l *Load, line string) []string {
	line = strings.TrimSpace(line)
	if strings.Count(line, ",") != 3 {
		l.Errors["-p"].GetErrors(3)
	}

	return strings.Split(line, ",")
}

func SaveInputs(l *Load, lineArr []string) {
	newPassword := create.PasswordInit(l.UserId)
	newPassword.Title = lineArr[0]
	newPassword.Url = lineArr[1]
	newPassword.Username = lineArr[2]
	newPassword.Password = lineArr[3]
	newPassword.CheckInputs()
	newPassword.Save()
}
