package main

import (
	"flag"
	"os"

	"github.com/peterbourgon/ff"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	SourceDir string
}

func MustConfig() *Config {
	fs := flag.NewFlagSet("checker", flag.ExitOnError)
	var (
		sourceDir = fs.String("d", ".", "directory with Dockerfiles to scan")
		debug     = fs.Bool("debug", false, "show debug logs")
	)
	ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.JSONParser),
		ff.WithEnvVarPrefix("DCHK"),
	)

	setupLogger(*debug)

	return &Config{
		SourceDir: *sourceDir,
	}
}

func setupLogger(debug bool) {
	log.SetFormatter(&log.JSONFormatter{})
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
