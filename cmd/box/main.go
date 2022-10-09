package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type flagConfig struct {
	debug bool
	mode  string
}

func main() {
	root := cobra.Command{
		Use:  "box",
		Long: "Data transfer plug-in framework",
	}

	cfg := flagConfig{}

	root.PersistentFlags().BoolVar(&cfg.debug, "debug", false, "enable debug log")
	root.PersistentFlags().StringVar(&cfg.mode, "mode", "normal", "specify the mode of service operation")

	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return nil
	}

	root.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println(cfg.debug)
		fmt.Println(cfg.mode)
		return nil
	}

	if err := root.Execute(); err != nil {
		log.Err(err).Send()
	}
}
