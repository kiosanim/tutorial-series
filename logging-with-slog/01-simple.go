// Este é um exemplo simples de logging
package main

import "log/slog"

func main() {
	slog.Info("Esta é uma mensagem de INFO")
	slog.Warn("Esta é uma mensagem de WARNING")
	/*
		Para que as mensagens de debug sejam apresentadas,
		o nível de log precisa estar setado para Debug
	*/
	slog.Debug("Esta é uma mensagem de DEBUG")
	slog.Error("Esta é uma mensagem de ERROR")
}
