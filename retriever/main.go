package main

type Retriever interface {
	Get(url string) string
	//interface中函数不用加function，本身就全都是函数了
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrievePoster interface {
	Retriever
	Poster
}

func session(s RetrievePoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	//系统接口
}
