package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	config := MustConfig()
	log.Debugf("Config: [%+v]", config)
}
