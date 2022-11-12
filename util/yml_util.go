package util

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadTransfer(path string, conf interface{}) {
	path = GetFullPath(path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, conf)
}
