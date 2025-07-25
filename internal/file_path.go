package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Setup() string {
	var fileName string
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to retreive current directory")
	}
	fmt.Printf("Where would you like your files to go? Type the number\n 1. Current Folder: %s\n 2. Desktop\n 3. Custom\n> ", dir)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "1":
		fileName = GetFileName()
		fmt.Printf("File Name: %s\n", fileName)
	case "2":
		//send to desktop
	case "3":
		fileName = GetFilePath()
		fmt.Printf("%s", fileName)
		//send to customer
	default:
		fmt.Print("Not an option, please select a number")
		return ""
	}
	setup := filepath.Join(dir, fileName)
	fmt.Printf("Setup file path is: %s", setup)
	return setup
}

func GetFilePath() string {
	fmt.Print("Enter full path:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dir := scanner.Text()

	fmt.Println("Using directory:", dir)
	return "string return"
}

func GetFileName() string {
	fmt.Print("Enter File Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dir := scanner.Text()
	if dir == "" {
		fmt.Print("Must enter file name\n")
		dir = GetFileName()
	}
	split := strings.Split(dir, " ")
	if len(split) > 1 {
		fmt.Print("File name can not be more than one word\n")
		dir = GetFileName()
	}
	return dir
}
