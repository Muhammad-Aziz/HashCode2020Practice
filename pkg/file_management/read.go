package file_management

import (
	"errors"
	"io/ioutil"
)

func Read(path string) (string, error) {
	str, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	// I'm expecting data from file
	if len(str) == 0 {
		return "", errors.New("Empty data set")
	}
	return string(str), nil
}
