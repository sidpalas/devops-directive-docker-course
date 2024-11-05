// Package jsontoml is a program that converts JSON to TOML.
//
// # Usage
//
// Reading from stdin:
//
//	cat file.json | jsontoml > file.toml
//
// Reading from a file:
//
//	jsontoml file.json > file.toml
//
// # Installation
//
// Using Go:
//
//	go install github.com/pelletier/go-toml/v2/cmd/jsontoml@latest
package main

import (
	"encoding/json"
	"io"

	"github.com/pelletier/go-toml/v2"
	"github.com/pelletier/go-toml/v2/internal/cli"
)

const usage = `jsontoml can be used in two ways:
Reading from stdin:
  cat file.json | jsontoml > file.toml

Reading from a file:
  jsontoml file.json > file.toml
`

func main() {
	p := cli.Program{
		Usage: usage,
		Fn:    convert,
	}
	p.Execute()
}

func convert(r io.Reader, w io.Writer) error {
	var v interface{}

	d := json.NewDecoder(r)
	err := d.Decode(&v)
	if err != nil {
		return err
	}

	e := toml.NewEncoder(w)
	return e.Encode(v)
}
