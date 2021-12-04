package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadFileToString(path string) string {
	b, _ := ioutil.ReadFile(path)
	return string(b)
}
