package main

import (
	"fmt"
	"github.com/prasan/db-go/approaches/naive"
	"log"
)

func main() {

	s := "Hello world again"
	dataString := []byte(s)

	err := naive.SaveData("C:\\Users\\prasa\\Downloads\\example.txt", dataString)
	if err != nil {
		log.Fatal(err)
	}

	err = naive.SaveData2("C:\\Users\\prasa\\Downloads\\example.txt", dataString)
	if err != nil {
		log.Fatal(err)
	}
	err = naive.SaveData3("C:\\Users\\prasa\\Downloads\\example.txt", dataString)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written .. ")

}
