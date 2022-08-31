package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not zip file")
	nonAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(nonAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
