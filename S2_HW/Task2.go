package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "log.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer file.Close()

	size, _ := file.Stat()

	buf := make([]byte, size.Size())

	_, err = file.Read(buf)
	if err != nil {
		fmt.Println("Невозможно прочитать последовательность байтов из файла!", err)
		return
	}
	fmt.Println(string(buf))

}
