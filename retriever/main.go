package main

import (
	"fmt"
	"languagethought/retriever/mock"
	real2 "languagethought/retriever/real"
)

type Retriever interface {
	Get(url string) string
	//interface中函数不用加function，本身就全都是函数了
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a"}
	fmt.Printf("%T %v\n", r, r)
	r = real2.Retriever{}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
}
