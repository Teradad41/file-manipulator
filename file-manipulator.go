package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type FileManipulator struct {
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (f FileManipulator) reveres(inputPath, outputPath string) error {
	inFile, err := os.Open(inputPath)
	handleError(err)
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	handleError(err)
	defer outFile.Close()

	buf := make([]byte, 1024)
	for {
		n, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		for i := n - 1; i >= 0; i-- {
			_, err := outFile.Write([]byte{buf[i]})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (f FileManipulator) copy(inputPath, outputPath string) error {
	inFile, err := os.Open(inputPath)
	handleError(err)
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	handleError(err)
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	handleError(err)

	return nil
}

func (f FileManipulator) duplicate(inputPath, n string) error {
	content, err := os.ReadFile(inputPath)
	handleError(err)

	var newContent []byte
	nInt, _ := strconv.Atoi(n)

	for i := 0; i < nInt; i++ {
		newContent = append(newContent, content...)
	}

	err = os.WriteFile(inputPath, newContent, 0644)
	handleError(err)

	return nil
}

func (f FileManipulator) replaceString(inputPath, needle, newString string) error {
	content, err := os.ReadFile(inputPath)
	handleError(err)

	str := string(content)
	replaced := strings.Replace(str, needle, newString, -1)
	bytesStr := []byte(replaced)

	err = os.WriteFile(inputPath, bytesStr, 0644)
	handleError(err)

	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: program <command> <inputFile> <arguments...>")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	f := FileManipulator{}

	switch command {
	case "reverse":
		handleError(f.reveres(args[0], args[1]))
	case "copy":
		handleError(f.copy(args[0], args[1]))
	case "duplicate":
		handleError(f.duplicate(args[0], args[1]))
	case "replace-string":
		handleError(f.replaceString(args[0], args[1], args[2]))
	default:
		fmt.Println("Invalid command.")
		os.Exit(1)
	}

	fmt.Println("Operation completed successfully.")
}
