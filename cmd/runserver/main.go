package main

import (
	"log"

	"olympics/pkg/commands/runserver"

	"github.com/spf13/cobra"
)

var l string

func main() {
	cmd := &cobra.Command{
		Use:   "runserver",
		Short: "configshop run server",
		RunE:  runserver.Run,
	}

	cmd.PersistentFlags().StringVar(&runserver.ConfigPath, "config", "", "config file (default is $HOME/.cobra.yaml)")
	if err := cmd.MarkPersistentFlagRequired("config"); err != nil {
		log.Fatal("MarkFlagRequired() failed")
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
