package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ImageEncode(r io.Reader, w io.Writer, ext string) error {
	m, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	switch ext {
	case ".jpg", ".jpeg":
		if err := png.Encode(w, m); err != nil {
			return err
		}
	case ".png":
		opts := &jpeg.Options{
			Quality: 100,
		}
		if err := jpeg.Encode(w, m, opts); err != nil {
			return err
		}
	default:
		return errors.New("specify jpeg or png for the image format.")
	}

	if err := png.Encode(w, m); err != nil {
		return err
	}

	return nil
}

func main() {
	root := "./test"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		r, err := os.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()

		out := path[:len(path)-len(filepath.Ext(path))] + ".png"
		w, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer w.Close()

		if err := ImageEncode(r, w, filepath.Ext(path)); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}
