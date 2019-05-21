// Package services contains the business logic and services
package services

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func containsAll(haystack []string, needles []string) bool {
	for _, n := range needles {
		if !contains(haystack, n) {
			return false
		}
	}
	return true
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " \n")
		if len(line) < 2 || strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func check(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	log.Print(s)
	return true
}
