package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("./ohsawa.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	writer, err := os.Create("./ohsawa.png")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	if err := png.Encode(writer, m); err != nil {
		log.Fatal(err)
	}

}
