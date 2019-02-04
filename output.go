package main

import (
	"io"
	"os"
)

type Outputter interface {
	Write(file string, r io.Reader) error
}

type InplaceOutputter struct {
}

func (InplaceOutputter) Write(file string, r io.Reader) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	fmt.Printf("Fixed %s\n", file)
	return err
}

type STDOUTOutputter struct {
}

func (STDOUTOutputter) Write(file string, r io.Reader) error {
	_, err := io.Copy(os.Stdout, r)
	return err
}
