package commands

import (
	"log/slog"

	"github.com/spf13/cobra"

	api "gororoba/internal"
	"gororoba/internal/config"
)

func NewServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "serve",
		Run: runServeCommand(),
	}

	return cmd
}

func runServeCommand() CommandFunction {
	return func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		if env == "" {
			env = config.DevelopmentEnv
		}

		slog.Info("Environment: " + env)
		appConfig := config.LoadConfig(env)

		api.Start(appConfig)
	}
}
