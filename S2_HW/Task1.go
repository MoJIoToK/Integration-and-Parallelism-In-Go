package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("S2_HW/log.txt")
	if err != nil {
		fmt.Println("Невозможно создать файл, ", err)
		return
	}
	defer file.Close()

	count := 1
	var answer string
	for {
		_, _ = fmt.Scan(&answer)
		if isExit(answer) {
			break
		}
		timeNow := time.Now()
		text := fmt.Sprintf("№%d. %v: %s\n", count, timeNow.Format("2006-01-02 15:04:05"), answer)
		file.WriteString(text)
		count++

	}
}

func isExit(answer string) bool {
	if answer == "exit" {
		return true
	}
	return false
}
