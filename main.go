package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func env(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}
