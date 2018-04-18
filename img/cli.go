// refer https://deeeet.com/writing/2014/12/18/golang-cli-test/
package img

import (
	"flag"
	"io"
	"os"
	"path/filepath"
)

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitError
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var dir, before, after string
	flags := flag.NewFlagSet("goimgconverter", flag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)
	flags.StringVar(&dir, "d", "", "path of conversion destination")
	flags.StringVar(&before, "b", "", "image extension before conversion")
	flags.StringVar(&after, "a", "", "image extension after conversion")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if err := cli.walk(dir); err != nil {
		return ExitError
	}

	return ExitCodeOK
}

func (cli *CLI) walk(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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

		if err := Convert(r, w, filepath.Ext(path)); err != nil {
			return err
		}

		return nil
	})
}
