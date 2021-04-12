package cmd

import (
	"main/server"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(uiCmd)
}

var uiCmd = &cobra.Command{
  Use:   "ui",
  Short: "Launch a web based UI tool that uses the CLI",
  Long:  `Additional long description here...`,
  Run: func(cmd *cobra.Command, args []string) {
    
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}
  },
}