package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	// CorrectAnswer is the goal.
	CorrectAnswer = "That's the right answer!"
	// AlreadySolved means our solution is already there.
	AlreadySolved = "Did you already complete it?"
	// WrongAnswer means the solution was incorrect.
	WrongAnswer = "That's not the right answer;"
	// WaitMore means you have to wait more.
	WaitMore = "You gave an answer too recently;"
	// NoSession means the provided token was not valid.
	NoSession = "To play, please identify yourself"
)

func correctOrNot(content string) (bool, error) {
	valid := []string{
		CorrectAnswer,
		AlreadySolved,
	}

	for _, check := range valid {
		if strings.Contains(content, check) {
			return true, nil
		}
	}

	if strings.Contains(content, NoSession) {
		return false, IncorrectAnswerError{
			Hint: NoSession,
			Wait: "",
		}
	}

	if strings.Contains(content, WrongAnswer) {
		re := regexp.MustCompile(
			`That's not the right answer; (?P<Hint>[^\.]+)..*(?P<Wait>Please wait [^.]+).`,
		)

		result := re.FindStringSubmatch(content)

		return false, IncorrectAnswerError{
			Hint: result[1],
			Wait: result[2],
		}
	}

	if strings.Contains(content, WaitMore) {
		re := regexp.MustCompile(
			`You gave an answer too recently; you .* again. (?P<Wait>[^\.]+).`,
		)

		result := re.FindStringSubmatch(content)

		return false, IncorrectAnswerError{
			Hint: "You gave an answer too recently",
			Wait: result[1],
		}
	}

	re := regexp.MustCompile(`<article><p>(?P<Content>.*)</p></article>`)
	result := re.FindStringSubmatch(content)

	return false, IncorrectAnswerError{
		Hint: result[1],
		Wait: "",
	}
}

// SubmitSolution downloads and saves the requested input file.
func (c *Client) SubmitSolution(year, day, part int, solution string) (bool, error) {
	form := url.Values{}
	form.Add("level", fmt.Sprintf("%d", part))
	form.Add("answer", solution)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day),
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return false, fmt.Errorf("unable to create request from context: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
		return false, fmt.Errorf("request error: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, DownloadError{StatusCode: resp.StatusCode}
	}

	content, _ := ioutil.ReadAll(resp.Body)

	return correctOrNot(string(content))
}
