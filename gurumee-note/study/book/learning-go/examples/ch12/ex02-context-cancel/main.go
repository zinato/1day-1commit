package main

import (
	"context"
	"learning-go/examples/ch12/ex02-context-cancel/pkg"
	"os"
)

func main() {
	ss := pkg.SlowServer()
	defer ss.Close()

	fs := pkg.FastServer()
	defer fs.Close()

	ctx := context.Background()
	pkg.CallBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
