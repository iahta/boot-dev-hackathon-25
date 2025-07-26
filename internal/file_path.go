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
	fmt.Printf("Where would you like your files to go? Type the number\n 1. Current Folder: %s\n 2. Custom\n> ", dir)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "1":
		fileName = GetFileName()
		fmt.Printf("File Name: %s\n", fileName)
	case "2":
		dir = GetFilePath()
	default:
		fmt.Print("Not an option, please select a number\n")
		return Setup()
	}
	setup := filepath.Join(dir, fileName)

	return setup
}

func GetFilePath() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter full path:")
		scanner.Scan()
		dir := scanner.Text()

		info, err := os.Stat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("File Path %s does not exist\n", dir)
			} else {
				fmt.Printf("Error getting file path info for %s", dir)
			}
			continue
		}
		if !info.IsDir() {
			fmt.Printf("%s is a file\n", dir)
			continue
		}

		fmt.Println("Using directory:", dir)
		return dir
	}

}

func GetFileName() string {
	fmt.Print("Enter File Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dir := scanner.Text()
	if dir == "" {
		fmt.Print("Must enter file name\n")
		return GetFileName()
	}
	split := strings.Split(dir, " ")
	if len(split) > 1 {
		fmt.Print("File name can not be more than one word\n")
		return GetFileName()
	}
	return dir
}
