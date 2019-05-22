package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/J-Rocke/gropro/algorithmus"
	"github.com/J-Rocke/gropro/ausgabe"
	"github.com/J-Rocke/gropro/eingabe"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("file or directory as commandline option mandatory")
		return
	}
	info, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var files []string
	if info.IsDir() {
		files, err = filepath.Glob(fmt.Sprintf("%s/*.txt", os.Args[1]))
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		files = []string{os.Args[1]}
	}
	for _, file := range files {
		if err := do(file); err != nil {
			fmt.Printf("error while handling: %s, error: %v", file, err)
		}
	}
}

func do(filename string) error {
	ac, err := eingabe.ParseFromFile(filename)
	if err != nil {
		return err
	}
	l := algorithmus.Loese(ac)
	return ausgabe.GNUPlotToFile(fmt.Sprintf("%s.gnu", strings.TrimSuffix(filename, ".txt")), l)
}
