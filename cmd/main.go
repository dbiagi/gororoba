package main

import (
	"gororoba/cmd/commands"
	"gororoba/internal/config"

	"github.com/spf13/cobra"
)

func main() {
	env := config.DevelopmentEnv

	rootCmd := rootCmd()
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "development", "Environment to run the application")
	rootCmd.AddCommand(commands.NewCreateRecipesCommand())
	rootCmd.AddCommand(commands.NewServeCommand())
	rootCmd.Execute()
}

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gororoba",
		Short: "Helper to gororoba application",
	}
}
