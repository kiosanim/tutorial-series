// Exemplo de como setar o nível default de logging
package main

import (
	"log/slog"
	"os"
)

func main() {
	options := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, options))
	logger.Info("Esta é uma mensagem de INFO")
	logger.Warn("Esta é uma mensagem de WARNING")
	logger.Debug("Esta é uma mensagem de DEBUG")
	logger.Error("Esta é uma mensagem de ERROR")
}
