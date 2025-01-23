package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "files//counter.txt"

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Error reading file : %v", err)
	}

	fmt.Println("File Content:")
	fmt.Println(string(content))
}
