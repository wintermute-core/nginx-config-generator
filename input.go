package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// paring input file

func parse(file string) Input {
	b, err := ioutil.ReadFile(file)
	check(err)

	i := Input{}
	err = yaml.Unmarshal(b, &i)
	check(err)
	return i
}
