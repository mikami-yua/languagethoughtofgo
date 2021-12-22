package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {
}

func (Retriever) Get(url string) string { //用不到形参就不写了，只写类型即可
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}