package main

import (
	"fmt"
	"log"
	"os"
	"math/rand"
	"time"
)

// 1. Persistence. How not to lose or corrupt your data. Recovering from a crash.
// 2. Indexing. Efficiently querying and manipulating your data. (B-tree).
// 3. Concurrency. How to handle multiple (large number of ) clients. And transactions.

// Naive approach
//persist data to a file
// problems : Concurrency : we are truncating the file before writing to it
// if someone else uses it the same time we might loose data .
//Writing data to files may not be atomic, depending on the size of the write. Con-
// current readers might get incomplete data.

// The data is probably still in the operating system’s page cache after the write syscall returns so the file state might be broken if te system
// crashes and reboots
func SaveData(path string, data []byte) error{
	fd, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0664)

	if err != nil {
		log.Fatal(err)
	}

	defer fd.Close()

	_, err = fd.Write(data)

	return err
}

// to solve some of these problems lets make the file writing processs atomic

//we will create a tmp file first and write our data in that then 
// when that is done we will rename our tmp file with the orignal file name
// after deleting the orignal file , Renaming a file in os is atomic(its either done or not done there is nothing in between these two states)

func SaveData2(path string, data []byte) error{

	rand.Seed(time.Now().UnixNano())

    // Generate a random integer between 0 and 99
    x := rand.Intn(100)

	tmpPath := fmt.Sprintf("%s.tmp.%d", path, x)

	fp, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)



}



func main(){

	s := "Hello world again";
	dataString := []byte(s)

	err := SaveData("C:\\Users\\prasa\\Downloads\\example.txt",dataString)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written .. ")

}