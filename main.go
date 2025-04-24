package main

import (
	"log/slog"

	"github.com/chaihaobo/boice-blog-api/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		slog.Error("application run failed:", slog.String("error", err.Error()))
	}
}
