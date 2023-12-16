package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal("Creating file error")
	}
	defer file.Close()

	msg := []byte("ТРЕТЬЬЬЬЬЯ ЛАБА????")
	_, err = file.Write(msg)
	if err != nil {
		log.Fatal("Cannot write to file")
	}
}
