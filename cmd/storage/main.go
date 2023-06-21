package main

import (
	"fmt"
	"log"

	"github.com/s-cod/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	file, err = st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it upload", file)

}
