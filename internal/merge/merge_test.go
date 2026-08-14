package merge_test

import (
	"testing"

	"github.com/LYH2263/go-config-layer/internal/merge"
	"github.com/LYH2263/go-config-layer/internal/model"
)

func TestMerge(t *testing.T) {
	m := merge.Merge(model.Map{"a": "1"}, model.Map{"a": "2", "b": "3"})
	if m["a"] != "2" || m["b"] != "3" {
		t.Fatalf("%v", m)
	}
}
