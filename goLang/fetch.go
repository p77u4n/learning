package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: coppying %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Fprintf(os.Stderr, "resp: %v", resp)
		//fmt.Printf("%s", b)
	}
}
