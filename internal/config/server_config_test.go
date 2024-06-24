package config_test

import (
	"encoding/json"
	"testing"

	"github.com/danielmoisa/trip-planner/internal/config"
)

func TestPrintServiceEnv(t *testing.T) {
	config := config.DefaultServiceConfigFromEnv()
	_, err := json.MarshalIndent(config, "", "  ")

	if err != nil {
		t.Fatal(err)
	}
}
