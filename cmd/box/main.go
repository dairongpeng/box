package main

import (
	"github.com/dairongpeng/box/log"
	"github.com/go-kit/log/level"
	"github.com/spf13/cobra"
	"os"
)

var (
	appName = "box"
)

type flagConfig struct {
	debug bool
	mode  string

	logConfig *log.Config
}

func main() {
	root := cobra.Command{
		Use:   "box",
		Short: appName,
		Long:  "Data transfer plug-in framework",
	}

	cfg := flagConfig{
		debug:     false,
		mode:      "normal",
		logConfig: &log.Config{},
	}

	root.PersistentFlags().BoolVar(&cfg.debug, "debug", false, "enable debug log. One of: [false, true]")
	root.PersistentFlags().StringVar(&cfg.mode, "mode", "normal", "specify the mode of service operation")
	root.PersistentFlags().StringVar(&cfg.logConfig.Level, "log.level", "info", "Only log messages with the given severity or above. One of: [debug, info, warn, error]")
	root.PersistentFlags().StringVar(&cfg.logConfig.Format, "log.format", "logfmt", "Output format of log messages. One of: [logfmt, json]")

	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return nil
	}

	root.RunE = func(cmd *cobra.Command, args []string) error {
		logger := log.New(cfg.logConfig)

		level.Info(logger).Log("debug", cfg.debug)
		level.Info(logger).Log("mode", cfg.mode)
		level.Info(logger).Log("format", cfg.logConfig.Format)
		return nil
	}

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
