package main

import (
	"main/server"

	"github.com/markbates/pkger"
	log "github.com/sirupsen/logrus"
)

func init() {
	 // set up logrus configuration
	 log.SetLevel(log.DebugLevel)
	 log.SetFormatter(&log.TextFormatter{ForceColors: true})
}

func main() {
	pkger.Include("/web/dist")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}