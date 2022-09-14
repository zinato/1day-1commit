package main

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func logic(ctx context.Context, src string) (string, error) {
	dst := strings.ToUpper(src)
	return dst, nil
}

func main() {
	ctx := context.Background()
	res, err := logic(ctx, "a string")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
