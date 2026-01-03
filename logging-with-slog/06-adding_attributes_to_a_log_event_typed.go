// Exemplo de como adicionar atributos a um evento de log forçando tipos de dados
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
		slog.String("service_name", "usuarios_job"),
		slog.String("operation", "importacao_de_arquivo"),
		slog.String("file_name", "usuarios.csv"),
		slog.Int("duration_ms", 250),
		slog.Int("number_of_lines_loaded", 333),
	)
}
