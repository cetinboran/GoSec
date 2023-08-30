package load

import (
	"bufio"
	"fmt"
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

	fmt.Println("Passwords saved successfully")
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

func Format(l *Load) {
	content := "title1,url1,username1,password1\ntitle2,url2,username2,password2\ntitle3,url3,username3,password3"

	file, err := os.Create("format.txt")
	if err != nil {
		l.Errors["--format"].GetErrors(1)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		l.Errors["--format"].GetErrors(2)
	}

	fmt.Println("Created a sample format txt file")
}
