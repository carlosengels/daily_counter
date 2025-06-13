package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// increaseRecursive ne recursively increments a by 1
func increaseRecursive(a int) int {
	if a == -1 {
		return 0 // Base case: when a is -1, return 0 (to ensure termination)
	}
	fmt.Println(a)
	return increaseRecursive(a-1) + 1 // Reduce a, then add 1 on the way back up
}

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

	// Modify counter
	updatedContent := strconv.Itoa(increaseRecursive(counterValue))

	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Error converting file content to integer: %v", err)
	}

	fmt.Println("Updated File Content:")
	fmt.Println(string(updatedContent))
}
