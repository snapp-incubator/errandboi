package cmd

import (
	"log"

	"github.com/snapp-incubator/errandboi/internal/cmd/server"
	"github.com/snapp-incubator/errandboi/internal/config"
	"github.com/snapp-incubator/errandboi/internal/logger"
	"github.com/spf13/cobra"
)

func Execute() {
	cfg := config.New()

	logg := logger.New(cfg.Logger)

	rootCmd := &cobra.Command{
		Use:   "errandboi",
		Short: "Give your errands to the errandboi!",
	}

	server.Register(rootCmd, cfg, logg)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error executing command")
	}
}
