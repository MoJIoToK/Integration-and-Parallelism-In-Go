package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	//подсоединение к серверу
	d, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		text, _, err := bufio.NewReader(os.Stdout).ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		if _, err = d.Write([]byte(text)); err != nil {
			log.Fatal(err)
		}

		uppetText := []byte{}
		if _, err := d.Read(uppetText); err != nil {
			log.Fatal(err)
		}
	}
}
