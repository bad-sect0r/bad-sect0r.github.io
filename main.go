package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	githubRawURL = "https://raw.githubusercontent.com/bad-sect0r/bad-sect0r/main/README.md"
)

func main() {
	e := echo.New()

	// Endpoint to serve README.md as raw text
	e.GET("/readme", func(c echo.Context) error {
		data, err := fetchReadme()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading README")
		}
		return c.String(http.StatusOK, string(data))
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// Fetch README.md from GitHub
func fetchReadme() ([]byte, error) {
	resp, err := http.Get(githubRawURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
