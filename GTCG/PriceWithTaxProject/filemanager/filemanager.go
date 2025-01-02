package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func ReadLines(fileName string) ([]string, error) {
	file, error := os.Open(fileName)

	if error != nil {
		file.Close()
		return nil, errors.New("failed to open file")
	}
	file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	error = scanner.Err()

	if error != nil {
		return nil, errors.New("an error occurred while reading a file")
	}

	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	error := encoder.Encode(data)

	if error != nil {
		return errors.New("failed to write to file")
	}

	defer file.Close()

	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}

}
