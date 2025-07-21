package modelstest

import (
	"fmt"
	"os"
	"testing"

	"github.com/Xe/project-template/internal"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func init() {
	internal.UnbreakDocker()
}

const (
	postgresDB       = "project"
	postgresUser     = "admin"
	postgresPassword = "hunter2"
)

func MaybeSpawnDB(t *testing.T) string {
	t.Helper()

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		if os.Getenv("USE_TEST_CONTAINERS") == "" {
			t.Skip("test requires test containers")
			return ""
		}

		testcontainers.SkipIfProviderIsNotHealthy(t)

		req := testcontainers.ContainerRequest{
			Image:      "postgres:16",
			WaitingFor: wait.ForLog("database system is ready to accept connections"),
			Env: map[string]string{
				"POSTGRES_DB":       postgresDB,
				"POSTGRES_USER":     postgresUser,
				"POSTGRES_PASSWORD": postgresPassword,
			},
		}
		postgresC, err := testcontainers.GenericContainer(t.Context(), testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		testcontainers.CleanupContainer(t, postgresC)
		if err != nil {
			t.Fatal(err)
		}

		containerIP, err := postgresC.ContainerIP(t.Context())
		if err != nil {
			t.Fatal(err)
		}

		dbURL = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", postgresUser, postgresPassword, containerIP, postgresDB)
	}

	return dbURL
}

func MaybeSpawnValkey(t *testing.T) string {
	redisURL := os.Getenv("REDIS_URL")

	if redisURL == "" {
		if os.Getenv("DONT_USE_NETWORK") != "" {
			t.Skip("test requires network egress")
			return ""
		}

		testcontainers.SkipIfProviderIsNotHealthy(t)

		req := testcontainers.ContainerRequest{
			Image:      "valkey/valkey:8",
			WaitingFor: wait.ForLog("Ready to accept connections"),
		}
		valkeyC, err := testcontainers.GenericContainer(t.Context(), testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		testcontainers.CleanupContainer(t, valkeyC)
		if err != nil {
			t.Fatal(err)
		}

		containerIP, err := valkeyC.ContainerIP(t.Context())
		if err != nil {
			t.Fatal(err)
		}

		return fmt.Sprintf("redis://%s:6379/0", containerIP)
	}

	return redisURL
}
