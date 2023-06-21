package main

import (
	"fmt"

	"github.com/s-cod/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello, go!", st)

}
