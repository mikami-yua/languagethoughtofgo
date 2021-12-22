package main

import (
	"fmt"
	"languagethought/testing"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
