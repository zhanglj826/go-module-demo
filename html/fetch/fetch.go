package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	return fmt.Sprintf("%s", b)
}
