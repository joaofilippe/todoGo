package infraDatabase

import (
	"os"
	"testing"
)

func Test_GetConfigFromEnv(t *testing.T){
	os.Setenv("DB_HOST", "12456")
	os.Setenv("DB_PORT", "6541")
	os.Setenv("DB_USER", "miau")
	os.Setenv("DB_PASSWORD", "miau12456")
	os.Setenv("DB_NAME", "miaudb")


	GetConfigFromEnv()
}