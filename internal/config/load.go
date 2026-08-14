package config

import (
	"github.com/LYH2263/go-config-layer/internal/merge"
	"github.com/LYH2263/go-config-layer/internal/model"
	"github.com/LYH2263/go-config-layer/internal/source"
)

type Config struct{ m model.Map }

func Load(file, env map[string]string) (*Config, error) {
	m := merge.Merge(source.FromFile(file), source.FromEnv(env))
	return &Config{m: m}, nil
}

func (c *Config) Get(k string) string { return c.m[k] }
