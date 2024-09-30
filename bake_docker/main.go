package main

import (
    "fmt"
    "os"
)

func main() {
	message := os.Getenv("HELLO_MSG")
    if message == "" {
        message = "Hello, World!"
    }

	fmt.Println(message)
}
