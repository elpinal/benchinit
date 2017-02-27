package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/elpinal/benchinit/pkg/benchinit"
)

func main() {
	err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: benchinit name")
	}

	err := os.Mkdir(args[0], 0777)
	if err != nil {
		return errors.Wrap(err, "run")
	}
	f, err := os.Create(filepath.Join(args[0], "main.go"))
	if err != nil {
		return errors.Wrap(err, "run")
	}
	defer f.Close()

	b := benchinit.New()
	s := b.Build()
	_, err = f.WriteString(s)
	if err != nil {
		return errors.Wrap(err, "run")
	}
	return nil
}