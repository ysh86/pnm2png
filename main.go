package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/spakin/netpbm"
)

func main() {
	fin, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	i, format, err := image.Decode(fin)
	fmt.Printf("Input format: %v\n", format)
	fmt.Printf("Image: %v, %#v\n", i.Bounds(), i.ColorModel())
	if err != nil {
		panic(err)
	}

	// conv netpbm.npcolor.RGBM to RGBA
	var img *image.RGBA
	switch i := i.(type) {
	case *netpbm.RGBM:
		img = image.NewRGBA(i.Bounds())
		for p := 0; p < len(img.Pix) / 4; p++ {
			img.Pix[p*4+0] = i.Pix[p*3+0]
			img.Pix[p*4+1] = i.Pix[p*3+1]
			img.Pix[p*4+2] = i.Pix[p*3+2]
			img.Pix[p*4+3] = 0xff
		}
	default:
		panic(fmt.Errorf("unknown format %T", i))
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
