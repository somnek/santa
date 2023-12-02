package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

const (
	INPUT_FILE = "input.txt"
)

func main() {
	// request data
	day := 1
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
	}

	// new request
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// cookies 🍪
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: x,
	}

	// client
	c := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	r.AddCookie(cookie)
	r.Header.Set("User-Agent", headers["User-Agent"])

	// send request
	response, err := c.Do(r)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("status code : %v", response.StatusCode)

	// read & save content
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	save(string(body))
}

func save(content string) error {
	f, err := os.Create(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}

	return nil
}
