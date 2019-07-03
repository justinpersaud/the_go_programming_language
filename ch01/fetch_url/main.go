// Fetch prints the content found at a URL.package main

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url_prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url_prefix, url) {
			fmt.Fprint(os.Stdout, "url doesn't have protocol, adding it in\n")
			url = url_prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "response code: %d\n", resp.StatusCode)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
	}
}
