package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() {
	filename := "/Users/gurumee92/Studies/1day-1commit/gurumee-note/study/book/learning-go/examples/ch11/ex02-io-02/my_data.txt.gz"
	r, closer, err := buildGZipReader(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer closer()

	counts, err := countLetters(r)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(counts)
}
