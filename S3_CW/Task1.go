package main

import (
	"fmt"
)

func main() {
	//1.var wg sync.WaitGroup
	//1. wg.Add(2)
	//в функции в качестве аргументов передают &wg

	firstChan := putBook()
	secondChan := deliveryBook(firstChan)
	thirdChan := burnBook(secondChan)

	fmt.Println(<-thirdChan)

	//1. go putBook(respChan)
	//1. go deliveryBook(respChan)
	//1. time.Sleep(6 * time.Millisecond)

	//1. wg.Wait()
	//1. burnBook()
}

func putBook() chan string {
	//1. time.Sleep(2 * time.Millisecond)
	//1. defer wg.Done()

	firstChan := make(chan string)
	go func() {
		firstChan <- "складываю книги"
	}()
	return firstChan
}

func deliveryBook(firstChan chan string) chan string {
	//1. time.Sleep(4 * time.Millisecond)
	//1. defer wg.Done()
	secondChan := make(chan string)
	fmt.Println(<-firstChan)
	go func() {
		secondChan <- "доставляю книги"
	}()
	return secondChan
}

func burnBook(secondChan chan string) chan string {
	thirdChan := make(chan string)
	fmt.Println(<-secondChan)
	go func() {
		thirdChan <- "сжигаю книги"
	}()
	return thirdChan
}
