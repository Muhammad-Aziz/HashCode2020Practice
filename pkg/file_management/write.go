package file_management

import "io/ioutil"

func Write(data, path string) error {
	d1 := []byte(data)
	return ioutil.WriteFile(path, d1, 0644)
}
