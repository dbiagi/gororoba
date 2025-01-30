package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	COMPOSE_PATH       = "../docker-compose.yml"
	COMPOSE_IDENTIFIER = "g_e2e"
)

func TestSetup(t *testing.T) {
	identifier := tc.StackIdentifier(COMPOSE_IDENTIFIER)
	files := tc.WithStackFiles(COMPOSE_PATH)
	compose, err := tc.NewDockerComposeWith(files, identifier)
	require.NoError(t, err, "Failed to create docker-compose")

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background()), tc.RemoveOrphans(true), tc.RemoveImagesLocal, "Failed to stop compose")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	err = compose.
		WaitForService("localstack", wait.ForExposedPort().WithStartupTimeout(30)).
		Up(ctx, tc.Wait(true))

	require.NoError(t, err, "Failed to start services")

	StartTestServer()
}
