// Package datafile read the files
package datafile

import (
	"bufio"
	"os"
)

// GetStrings returns list of strings from file
func GetStrings(fileName string) ([]string, error) {
	var stringList []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringList = append(stringList, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return stringList, nil
}
