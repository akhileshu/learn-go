package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// "os"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Elodie")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

const englishHelloPrefix = "Hello, "

// Greet writes a greeting to the writer
func Greet(writer io.Writer, name string) {
	/*
		buffer.WriteString(englishHelloPrefix)
		buffer.WriteString(s)
	*/
	fmt.Fprintf(writer, "Hello, %s", name)
}
