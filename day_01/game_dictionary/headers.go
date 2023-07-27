package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// response body
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

// request body
type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserId    string `json:"user_id"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, `usage: simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}

// 查询方法
func query(word string) {
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"zh2en","source":"你好"}`)
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.1 Safari/605.1.15")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Content-Length", "40")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Host", "api.interpreter.caiyunai.com")
	req.Header.Set("version", "1.8.0")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("device-id", "9390caadab4f151bfe472f48ab1a8921")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 判断正确请求
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	//fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dictResponse)

	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}
