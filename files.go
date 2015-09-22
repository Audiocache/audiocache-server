package main

import (
	"encoding/base64"
	"io/ioutil"
)

func writeFile(postcache PostCache, filename string) error {
	data := postcache.Data
	var binary []byte

	n, err := base64.StdEncoding.Decode(binary, []byte(data))
	if n == 0 || err != nil {
		return err
	}

	err = ioutil.WriteFile("files/"+filename, binary, 0644)
	if err != nil {
		return err
	}

	return nil
}
