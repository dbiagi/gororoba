package commands

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"gororoba/internal/config"
	"gororoba/internal/controller"
	"gororoba/internal/handler"
	"gororoba/internal/repository"

	"github.com/aws/aws-sdk-go/service/dynamodb"
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
		startTime := time.Now()

		env, _ := cmd.Flags().GetString("env")
		if env == "" {
			env = config.DevelopmentEnv
		}

		slog.Info("Environment: " + env)
		appConfig := config.LoadConfig(env)

		config.ConfigureLogger(appConfig.AppConfig)

		slog.Info("Connecting to dynamoDB ....")
		dynamoDB := connectToDynamoDB(appConfig.AWSConfig)

		slog.Info("Starting server ....")
		srv, router := createServer(appConfig.WebConfig)

		slog.Info("Creating resources ....")
		appResources := createControllers(dynamoDB)

		slog.Info("Registering routes and serving ....")
		registerRoutesAndServe(router, appResources)

		slog.Info(fmt.Sprintf("Application ready. Time elapsed: %v", time.Since(startTime)))

		slog.Info("Configuring graceful shutdown.")
		configureGracefullShutdown(srv, appConfig.WebConfig)
	}
}

func connectToDynamoDB(awsConfig config.AWSConfig) *dynamodb.DynamoDB {
	dynamoDB, err := config.CreateDynamoDBConnection(awsConfig)
	if err != nil {
		slog.Error("Error connecting to dynamodb.", slog.String("error", err.Message))
		panic(err)
	}

	return dynamoDB
}

func createServer(webConfig config.WebConfig) (*http.Server, *mux.Router) {
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:         ":" + webConfig.Port,
		Handler:      router,
		IdleTimeout:  webConfig.IdleTimeout,
		ReadTimeout:  webConfig.ReadTimeout,
		WriteTimeout: webConfig.WriteTimeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			slog.Error("Error starting server.", slog.String("error", err.Error()))
		}
	}()

	return srv, router
}

func createControllers(db *dynamodb.DynamoDB) Controllers {
	recipeRepository := repository.NewRecipeRepository(db)
	healthCheckHandler := handler.NewHealthCheckHandler()
	suggestionHandler := handler.NewSuggestionHandler()
	recipeHandler := handler.NewRecipesHandler(recipeRepository, suggestionHandler)

	return Controllers{
		RecipesController:     controller.NewRecipesController(recipeHandler),
		HealthCheckController: controller.NewHealthCheckController(healthCheckHandler),
	}
}

func registerRoutesAndServe(router *mux.Router, controllers Controllers) {
	router.Use(config.TraceIdMiddleware)
	router.HandleFunc("/health", controller.HandleRequest(controllers.HealthCheckController.Check)).Methods("GET")
	router.HandleFunc("/health/complete", controller.HandleRequest(controllers.HealthCheckController.CheckComplete)).Methods("GET")
	router.HandleFunc("/recipes/by-category", controller.HandleRequest(controllers.RecipesController.GetRecipesByCategory)).Methods("GET")
	router.HandleFunc("/recipes/suggestion", controller.HandleRequest(controllers.RecipesController.GetSuggestion)).Methods("GET")
}

func configureGracefullShutdown(server *http.Server, webConfig config.WebConfig) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), webConfig.ShutdownTimeout)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("Shutting down server")
	os.Exit(0)
}
