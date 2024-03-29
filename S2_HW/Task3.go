package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "S2_HW/FILE_CHMOD.txt"
	file, err := os.Create(fileName)
	if err := os.Chmod(fileName, 0444); err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	writer.WriteString("Say hi!")
	writer.WriteString("\n")
	writer.WriteRune('a')
	writer.WriteString("\n")
	writer.WriteByte(67) //C
	writer.WriteString("\n")
	writer.Write([]byte{65, 66, 67}) //A,B,C
	writer.WriteString("\n")
	writer.Flush()

	file.Close()

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println(err)
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
