// Este é um exemplo do uso de logging com saída no formato TEXTO
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Esta é uma mensagem de INFO")
	logger.Warn("Esta é uma mensagem de WARNING")
	/*
		Para que as mensagens de debug sejam apresentadas,
		o nível de log precisa estar setado para Debug
	*/
	logger.Debug("Esta é uma mensagem de DEBUG")
	logger.Error("Esta é uma mensagem de ERROR")
}
