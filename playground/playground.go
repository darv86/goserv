package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("https://feeds.megaphone.fm/strike-force-five")
	if err != nil {
		fmt.Println(err.Error())
	}
	body := resp.Body
	defer body.Close()
	data := make([]byte, 0)
	chunk := make([]byte, 1024)
	for {
		n, err := body.Read(chunk)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
		}
		if n == 0 || err == io.EOF {
			fmt.Println("nothing to read")
			break
		}
		data = append(data, chunk[:n]...)
	}
	fmt.Println(string(data))
}
