package main

import (
	"io"
)

func main() {
	var w io.Writer

	w = nil
	w.Write([]byte("a"))
}
