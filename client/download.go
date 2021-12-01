package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// DownloadInput downloads the requested input file.
func (c *Client) DownloadInput(year, day int) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request from context: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:       "session",
		Value:      c.SessionToken,
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	})

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, DownloadError{StatusCode: resp.StatusCode}
	}

	return resp.Body, nil
}

// DownloadAndSaveInput downloads and saves the requested input file.
func (c *Client) DownloadAndSaveInput(year, day int, targetFile string) error {
	file, err := c.DownloadInput(year, day)
	if err != nil {
		return err
	}

	defer file.Close()

	outputFile, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("unable to create output file: %w", err)
	}

	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		return fmt.Errorf("unable to copy response into the output file: %w", err)
	}

	return nil
}
