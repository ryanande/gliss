package main

import (
	"main/cmd"

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
	cmd.Execute()
}
