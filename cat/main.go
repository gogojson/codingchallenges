package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

var (
	runMode  = ""
	withLine = false
)

func main() {
	// This program supports both Args and standard in, but needs at lest one
	// Check if there are Args
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			if v != "-n" && v != "-" {
				slog.Info("there are input arguments set to args mode", slog.Any("os.Args", v))
				runMode = "ARG"
			} else {
				withLine = true
			}
		}
	}
	// Check if there are Stdin
	s, err := os.Stdin.Stat()
	if err != err {
		slog.Error("error while getting stdin stat")
		panic(err.Error())
	}
	if (s.Mode() & os.ModeCharDevice) == 0 {
		slog.Info("Stdin is set")
		runMode = "STDIN"
	}

	switch {
	case runMode == "ARG":
		r, err := args(os.Args[1:])
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to read arg file %s", err.Error()))
			return
		}
		fmt.Println("---------------------------------")
		if withLine {
			fmt.Println(numberLine(r))
		} else {
			fmt.Println(r)
		}
	case runMode == "STDIN":
		r := os.Stdin
		b, err := io.ReadAll(r)
		if err != nil {
			slog.Error("error while reading")
			panic(err.Error())
		}
		fmt.Println("---------------------------------")
		if withLine {
			fmt.Println(numberLine(string(b)))
		} else {
			fmt.Println(string(b))
		}

		return
	default:
		slog.Error("Need to provide either arguments or standard in in the terminal")
		return
	}
}

func stdin() {

}

func args(a []string) (string, error) {

	result := ""
	// Check the number of args
	for _, v := range a {
		// Check if v is a existing path
		if v == "-n" || v == "-" {
			continue
		}
		b, err := os.ReadFile(v)
		if err != nil {
			return "", err
		}
		result += string(b)
	}
	return result, nil

}

func numberLine(s string) string {
	ss := strings.Split(s, "\n")
	result := ""
	for i, r := range ss {
		r = strings.TrimSpace(r)
		if r == "" && i == len(ss)-1 {
			continue
		}
		result += fmt.Sprintf("%d. %s\n", i, r)
	}
	return result

}
