package commands

import (
	"github.com/spf13/cobra"
)

type CommandFunction func(*cobra.Command, []string)
