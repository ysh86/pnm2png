package main

import (
	"fmt"
	"os"
	"image"
	"image/png"
	_ "github.com/spakin/netpbm"
)

func main() {
	fin, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	img, format, err := image.Decode(fin)
	fmt.Printf("Input format: %v\n", format)
	if err != nil {
		panic(err)
	}

	fout, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	err = png.Encode(fout, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}