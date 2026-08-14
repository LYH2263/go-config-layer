package main

import (
	"fmt"

	"github.com/LYH2263/go-config-layer/internal/config"
)

func main() {
	c, _ := config.Load(map[string]string{"addr": ":8080"}, map[string]string{"addr": ":9090"})
	fmt.Println(c.Get("addr"))
}
