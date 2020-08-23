package main

import (
	"log"
	"os"
)

// SaveMessage sends parsed data to a file
func SaveMessage(s string, filePath string) bool {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write([]byte(s)); err != nil {
		return false
	}
	return true
}