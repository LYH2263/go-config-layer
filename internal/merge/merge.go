package merge

import "github.com/LYH2263/go-config-layer/internal/model"


// Merge：env 覆盖 file。
func Merge(file, env model.Map) model.Map {
	out := model.Map{}
	for k, v := range file {
		out[k] = v
	}
	for k, v := range env {
		out[k] = v
	}
	return out
}

