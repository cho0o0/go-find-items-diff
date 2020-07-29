package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var first string
	if len(os.Args) <= 1 {
		first = "first.txt"
	} else {
		first = os.Args[1]
	}

	var second string
	if len(os.Args) <= 2 {
		second = "second.txt"
	} else {
		second = os.Args[2]
	}

	firstItems, firstErr := readLines(first)
	if firstErr != nil {
		panic(fmt.Sprintf("%s not readable.", first))
	}
	secondItems, secondErr := readLines(second)
	if secondErr != nil {
		panic(fmt.Sprintf("%s not readable.", second))
	}

	result := make(map[string]string)
	
	Outer: for _, firstItem := range firstItems {
		if _, ok := result[firstItem]; ok {
			continue
		}
		for _, secondItem := range secondItems {
			if strings.EqualFold(firstItem, secondItem) {
				result[firstItem] = "both"
				continue Outer
			}
		}
		result[firstItem] = first + " only"
	}

	for _, secondItem := range secondItems {
		if _, ok := result[secondItem]; ok {
			continue
		}
		result[secondItem] = second + " only"
	}

	for k, v := range result {
		println(fmt.Sprintf("%s: %s", k, v))
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
