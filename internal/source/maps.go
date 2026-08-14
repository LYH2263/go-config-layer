package source

import "github.com/LYH2263/go-config-layer/internal/model"

func FromFile(m map[string]string) model.Map { return model.Map(m) }
func FromEnv(m map[string]string) model.Map  { return model.Map(m) }
