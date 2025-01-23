package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := "files//counter.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file : %v", err)
	}

	counterValue, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatalf("Error converting file content to integer: %v", err)
	}

	counterValue++

	updatedContent := strconv.Itoa(counterValue)

	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Error converting file content to integer: %v", err)
	}

	fmt.Println("Updated File Content:")
	fmt.Println(string(updatedContent))
}
