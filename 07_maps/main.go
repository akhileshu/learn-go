package main

import "errors"

func main() {
	// entry point
}

type Dictionary map[string]string

var errNotFound = errors.New("could not find the word you were looking for")
func (dictionary Dictionary) Search(word string) (string,error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	 return  definition, nil
}

