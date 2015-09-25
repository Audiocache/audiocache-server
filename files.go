package main

import (
	"encoding/base64"
	"io/ioutil"
)

func writeFile(postcache PostCache, filename string) error {
	data := postcache.Data
	var binary []byte

	binary, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(config.API.Files+filename, binary, 0644)
	if err != nil {
		return err
	}

	return nil
}
