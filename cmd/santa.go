package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	USER_AGENT = "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0"
)

func download(cmd *cobra.Command, day string) {
	// args
	session := viper.GetString("aoc_session")
	year := 2023
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%s/input", year, day)
	headers := map[string]string{
		"User-Agent": USER_AGENT,
	}

	_, body := request(url, session, headers)
	fmt.Print(string(body))
}

func request(url, session string, headers map[string]string) (int, []byte) {
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
		Value: session,
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

	// read & save content
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return response.StatusCode, body

}
