package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	conf "github.com/rwxrob/conf-go"
)

func main() {
	checked := conf.NewMap()

	err := filepath.Walk("/", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("Failed to access %s\n", path)
			return err
		}

		if checked.Get(path) == "" {
			unixtimestamp := fmt.Sprintf("%d", info.ModTime().Unix())

			// checked.Set("/dev/null", 1640377491)
			checked.Set(path, unixtimestamp)
		}

		control := checked.Get(path)
		log.Printf("%s: %s\n", path, control)
		if fmt.Sprintf("%d", info.ModTime().Unix()) != control {
			log.Printf("Found change at %s\n", path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
