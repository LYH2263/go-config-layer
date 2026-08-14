package merge

import "github.com/LYH2263/go-config-layer/internal/model"


// Merge：本意 env 覆盖 file，实现写反了。
func Merge(file, env model.Map) model.Map {
	out := model.Map{}
	for k, v := range env {
		out[k] = v
	}
	for k, v := range file {
		out[k] = v
	}
	return out
}

