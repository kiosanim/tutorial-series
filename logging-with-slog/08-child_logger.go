// Exemplo de child logging
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	child := logger.With(
		slog.Group("service_info",
			slog.String("service_name", "usuarios_job"),
			slog.String("operation", "importacao_de_arquivo"),
		),
	)
	child.Info("details", slog.String("status", "inicializado"))
	child.Info("details", slog.String("status", "finalizado"))
}
