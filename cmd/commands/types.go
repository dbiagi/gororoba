package commands

import (
	"gororoba/internal/controller"

	"github.com/spf13/cobra"
)

type CommandFunction func(*cobra.Command, []string)

type Controllers struct {
	controller.HealthCheckController
	controller.RecipesController
}
