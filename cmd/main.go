package main

import (
	"log/slog"

	"github.com/JWindy92/go_app_utils/pkg/logging"
)

func main() {
	logging.InitLogger()
	slog.Info("Logger initialized")
}
