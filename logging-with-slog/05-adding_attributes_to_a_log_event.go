// Exemplo de como adicionar atributos a um evento de log
package main

import (
	"log/slog"
	"os"
)

func main() {
	options := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, options))
	logger.Info(
		"Mensagem de INFO",
		"service_name", "usuarios_job",
		"operation", "importacao_de_arquivo",
		"file_name", "usuarios.csv",
		"duration_ms", 250,
		"number_of_lines_loaded", 333,
	)
}
