package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	//Открытие соединения
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server is running!")

	//Чтобы можно было принимать соединение с сокетом, открытое ранее
	con, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(con)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Line received: ", line)

		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatal(err)
		}
	}
}
