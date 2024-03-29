package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer
	fileName := "S2_HW/log.txt"
	//file, err := os.Create("S2_HW/log.txt")
	//if err != nil {
	//	fmt.Println("Невозможно создать файл, ", err)
	//	return
	//}
	//defer file.Close()

	count := 1
	var answer string
	for {
		_, _ = fmt.Scan(&answer)
		if isExit1(answer) {
			break
		}
		timeNow := time.Now()
		text := fmt.Sprintf("№%d. %v: %s\n", count, timeNow.Format("2006-01-02 15:04:05"), answer)
		b.WriteString(text)
		count++

	}

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохранненый лог: ")
	fmt.Println(string(resultBytes))

}

func isExit1(answer string) bool {
	if answer == "exit" {
		return true
	}
	return false
}
