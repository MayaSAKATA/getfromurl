// utils/httputils.go
package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CheckStatus(code int) {
	var message string

	switch {
	case code >= 100 && code <= 199:
		message = "Informational response"
	case code >= 200 && code <= 299:
		message = "Successful response"
	case code >= 300 && code <= 399:
		message = "Redirection message"
	case code >= 400 && code <= 499:
		message = "Client error message"
	case code >= 500 && code <= 599:
		message = "Server error response"
	default:
		message = "Unknown status"
	}
	statusText := http.StatusText(code)
	fmt.Printf("%s\nStatus Code: %d %s\n", message, code, statusText)
}

func CheckContent(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading http response: %s\n", err)
	}
	return string(body)
}

func WriteInFile(fileName string, content string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}
