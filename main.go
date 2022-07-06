package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("httpclient <file>")
		return
	}

	proto, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	req, err := http.ReadRequest(bufio.NewReader(proto))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	content, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println(string(content))
}
