package main

import (
	"io"
	"log/slog"
	"os"
)

func main() {
	var r io.Reader
	f := "cat/test.txt"

	if len(os.Args) == 2 {
		f = os.Args[1]
	}
	slog.Info("all flags parsed", slog.String("file", f))

	//TODO: check if file exists
	if _, err := os.Stat(f); os.IsNotExist(err) {
		slog.Error(err.Error())
		return
	}
	// b, err := os.ReadFile(f)
	// if err != err {
	// 	slog.Error("error while reading file")
	// 	panic(err.Error())
	// }
	// slog.Info(string(b))

	r = os.Stdin

	b, err := io.ReadAll(r)
	if err != nil {
		slog.Error("error while reading")
		panic(err.Error())
	}
	slog.Info(string(b))
}
