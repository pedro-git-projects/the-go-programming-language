package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
// buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
func FetchWCopy() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix.
func checkPrefix(url *string) {
	builder := strings.Builder{}
	const prefix = "http://"
	if !strings.HasPrefix(*url, prefix) {
		builder.Write([]byte(prefix))
		builder.Write([]byte(*url))
		*url = builder.String()
	}
}

func FetchWPrefix() {
	for _, url := range os.Args[1:] {
		checkPrefix(&url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
func FetchWStatus() {
	for _, url := range os.Args[1:] {
		checkPrefix(&url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "Status: %s\n", resp.Status)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
