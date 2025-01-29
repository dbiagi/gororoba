package commands

import (
	"encoding/json"
	"fmt"
	"gororoba/internal/config"
	"gororoba/internal/domain"
	"gororoba/internal/handler"
	"gororoba/internal/repository"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func NewCreateRecipesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "create-recipes",
		Run:       runCreateRecipeCommand(),
		ValidArgs: []string{"file-path"},
	}

	return cmd
}

func runCreateRecipeCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		appConfig := config.LoadConfig(env)
		config.ConfigureLogger(appConfig.AppConfig)

		dynamoDb, dynamoConnectError := config.CreateDynamoDBConnection(appConfig.AWSConfig)

		if dynamoConnectError != nil {
			slog.Error(fmt.Sprintf("Failed to run command %v.", dynamoConnectError))
			return
		}
		r := repository.NewRecipeRepository(dynamoDb)
		h := handler.NewRecipesHandler(r, nil)

		jsonFile := args[0]

		slog.Info(fmt.Sprintf("Creating recipes from file %s", jsonFile))

		bytes, readError := os.ReadFile(jsonFile)
		if readError != nil {
			slog.Error(fmt.Sprintf("Failed to read file: %v", readError))
			return
		}

		var recipes []domain.Recipe
		if err := json.Unmarshal(bytes, &recipes); err != nil {
			slog.Error(fmt.Sprintf("Failed to unmarshal JSON: %v", err))
			return
		}

		for _, recipe := range recipes {
			insertRecipe(&recipe, h)
		}
	}
}

func insertRecipe(recipe *domain.Recipe, h handler.RecipesHandler) {
	slog.Info(fmt.Sprintf("Creating recipe: %s", recipe.Title))
	h.CreateRecipe(recipe)
}
