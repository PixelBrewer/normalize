// package config is to ingest and set config values
package config

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func ingestConfigToml() string {
	// This will always be the toml file name.
	// tomlFileName := "normalize.toml"
	return ""
}

func FindConfigToml() string {
	root := "/home/alex/"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(root)
		return nil
	})
	return ""
}
