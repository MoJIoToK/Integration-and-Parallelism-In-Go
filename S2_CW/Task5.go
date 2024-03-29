package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("S2_CW/some.txt")
	if err := os.Chmod("S2_CW/some.txt", 0444); err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer.WriteString("Say hi!")
	writer.WriteString("\n")
	writer.WriteRune('a')
	writer.WriteString("\n")
	writer.WriteByte(67) //C
	writer.WriteString("\n")
	writer.Write([]byte{65, 66, 67}) //A,B,C
	writer.WriteString("\n")
	writer.Flush()
}
