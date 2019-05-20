package services

import (
	"bufio"
	"github.com/fsnotify/fsnotify"
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

func watch(path string, mod chan string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				//log.Printf("event %v", event)
				if !ok {
					mod <- "error:"
					return
				}
				if event.Op & fsnotify.Write == fsnotify.Write {
					mod <- "modified:" + event.Name
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					mod <- "error:" + err.Error()
					return
				}
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func check(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	log.Print(s)
	return true
}
