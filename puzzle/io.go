package puzzle

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func inputAsReader(inputValue, inputFile string) (io.ReadCloser, error) {
	var (
		input io.ReadCloser
		err   error
	)

	if inputFile != "" {
		input, err = os.Open(inputFile)
		if err != nil {
			return nil, fmt.Errorf("unable to open file: %w", err)
		}
	} else {
		input = ioutil.NopCloser(strings.NewReader(inputValue))
	}

	return input, nil
}
