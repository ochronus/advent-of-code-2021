package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
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

func ReadFileToIntList(path string) []int {

	s := ReadFileToString(path)
	i := strings.Split(s, ",")
	list := []int{}
	for _, elem := range i {
		list = append(list, StrToInt(elem))
	}
	return list
}
