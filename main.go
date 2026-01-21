package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
func SaveData(path string, data []byte) error {
	fd, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

	if err != nil {
		log.Fatal(err)
	}

	defer fd.Close()

	_, err = fd.Write(data)

	return err
}

// to solve some of these problems lets make the file writing process atomic

//we will create a tmp file first and write our data in that then
// when that is done we will rename our tmp file with the orignal file name
// after deleting the orignal file , Renaming a file in unix based os is atomic ( ie its either done or not done there is nothing in between these two states)

//In this new implementation we fixed one thing ie Made our data updation atomic now
// when we update our file we have either of the two states File updated or FIle not updated
//There is nothing in between

func SaveData2(path string, data []byte) error {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randomFileName := r.Intn(100)

	tmpPath := fmt.Sprintf("%s.tmp.%d", path, randomFileName)

	fp, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()

	_, err = fp.Write(data)

	if err != nil {
		os.Remove(tmpPath)
		log.Fatal(err)
	}

	return os.Rename(tmpPath, path)
}

// THis solves out atomicity problem but we still have a problem we do not when the data
// will be written to the disk (persistence problem) as when we do a write to the disk
// the data will still be kept in mem for sometime the cpu optimizes this process
// therefore we need some predicable behaviour  // Enterns fsync and fdatasync
// fsync is a way to tell the cpu to flush the data to the disk


func SaveData3(path string, data []byte) error {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randomFileName := r.Intn(100)

	tmpPath := fmt.Sprintf("%s.tmp.%d", path, randomFileName)

	fp, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmpPath)
		log.Fatal(err)
	}
	
	//All we do is call the fsync after writing the data to the file
	// this essures any file buffers are flushed to disk

	err = fp.Sync()
	if err != nil {
		os.Remove(tmpPath)
		log.Fatal(err)
	}

	return os.Rename(tmpPath, path)
}


func main() {

	s := "Hello world again"
	dataString := []byte(s)

	err := SaveData("C:\\Users\\prasa\\Downloads\\example.txt", dataString)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written .. ")

}
