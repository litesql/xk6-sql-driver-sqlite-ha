package sqliteha_test

import (
	"context"
	_ "embed"
	"fmt"
	"runtime"
	"testing"
	"time"

	//"github.com/grafana/xk6-sql/sqltest"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

//go:embed testdata/script.js
var script string

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	if runtime.GOOS != "linux" {
		t.Skip("Works only on Linux (Testcontainers)")
	}

	ctx := context.Background()

	haContainer, err := testcontainers.Run(
		ctx,
		"ghcr.io/litesql/ha:latest",
		testcontainers.WithExposedPorts("8080/tcp"),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("8080/tcp").WithStartupTimeout(10*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start container")

	testcontainers.CleanupContainer(t, haContainer)

	host, err := haContainer.Host(ctx)
	require.NoError(t, err)

	mappedPort, err := haContainer.MappedPort(ctx, "8080/tcp")
	require.NoError(t, err)

	connectionString := fmt.Sprintf("http://%s:%s", host, mappedPort.Port())
	t.Logf("HA is ready and listening at: %s", connectionString)

	//sqltest.RunScript(t, "remote-ha", connectionString, script)
}
