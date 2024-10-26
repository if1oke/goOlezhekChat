package main

import (
	"fmt"
	"github.com/if1oke/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello world"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("uploaded", file)

	foundFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatalf("File with ID %s not found", file.ID)
	}
	fmt.Println("found", foundFile)

	fmt.Println(st.Files)
	// fmt.Println(st.filesBackup) // Cause Error
}
