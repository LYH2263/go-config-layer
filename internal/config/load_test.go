package config_test

import (
	"testing"

	"github.com/LYH2263/go-config-layer/internal/config"
)

func TestEnvOverridesFile(t *testing.T) {
	c, err := config.Load(map[string]string{"addr": ":8080"}, map[string]string{"addr": ":9090"})
	if err != nil {
		t.Fatal(err)
	}
	if c.Get("addr") != ":9090" {
		t.Fatalf("got %q", c.Get("addr"))
	}
}
