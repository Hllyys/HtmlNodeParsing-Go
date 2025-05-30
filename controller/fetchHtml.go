package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// FetchHtml: verilen URL'den HTTP GET ile HTML içeriği çeker
func FetchHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		checkError(err)
	}()

	var builder strings.Builder
	_, err = io.Copy(&builder, resp.Body)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
