package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := MustConfig()
	log.Debugf("Config: [%+v]", config)

	filepaths := make([]string, 0, 10)
	filepath.Walk(config.SourceDir, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(strings.ToLower(info.Name()), "docker") {
			filepaths = append(filepaths, path)
		}

		return nil
	})

	for _, path := range filepaths {
		file, err := os.Open(path)
		if err != nil {
			continue
		}

		log.Debugf("Parse file: [%s]", path)
		from := getFROMValue(file)
		if from == "" {
			continue
		}

		log.WithField("FROM", from).Infof("Found FROM command")
	}
}

func getFROMValue(file io.Reader) string {
	res, err := parser.Parse(file)
	if err != nil {
		log.WithError(err).Debugf("Failed to parse file")
		return ""
	}

	for _, child := range res.AST.Children {
		if child.Value != "from" {
			log.Debugf("Skipping command: [%s]", child.Value)
			continue
		}

		value := ""
		for n := child.Next; n != nil; n = n.Next {
			value += n.Value
		}

		return value
	}

	log.Debugf("Couldn't file a FROM commang")
	return ""
}
