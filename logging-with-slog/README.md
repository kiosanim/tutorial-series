# logging-with-slog
Autor: Fábio Sartori

Exemplos de logging em golang com a biblioteca slog

## Níveis de Log

Os níveis padrões de log, representam a importância / severidade de um evento de log.
Quanto maior o valor associado, mais importante / severo é o evento.

| Nível | Descrição / Uso                              | Valor Associado |
|-------|----------------------------------------------|-----------------|
| DEBUG | Serve para detalhes técnicos usados em debug | -4              |
| INFO  | Registrar operações normais do sistema       | 0               |
| WARN  | Indicar situações anormais sem falha         | 4               |
| ERROR | Erros que afetam o funcionamento esperado    | 8               |

Quando determinado nível de log está setado, todos eventos associados ao valor dele e os maiores são apresentados.

Ex:

| Nível | Eventos Apresentados     |
|-------|--------------------------|
| DEBUG | DEBUG, INFO, WARN, ERROR |
| INFO  | INFO, WARN, ERROR        |
| WARN  | WARN, ERROR              |
| ERROR | ERROR                    |

---

**COMO EXECUTAR OS EXEMPLOS**

```bash
# go run <NOME_DO_PROGRAMA>
# Ex:
go run 01-simple.go
```

## Exemplo Simples de Logging

```go
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
```

---

## Log no formato JSON

```go
// Este é um exemplo do uso de logging com saída no formato JSON
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Esta é uma mensagem de INFO")
	logger.Warn("Esta é uma mensagem de WARNING")
	/*
		Para que as mensagens de debug sejam apresentadas,
		o nível de log precisa estar setado para Debug
	*/
	logger.Debug("Esta é uma mensagem de DEBUG")
	logger.Error("Esta é uma mensagem de ERROR")
}
```

---

## Log no formato TEXTO

```go
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
```

---

## Setando o nível padrão de log

```go
// Exemplo de como setar o nível default de log
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
```

---

## Adicionando atributos a um evento de log

```go
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
```

---

## Adicionando atributos a um evento de log, tipando-os

```go
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
```

## Agrupando eventos de log

```go
// Exemplo de como agrupar atributos em um evento de log
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
		slog.Group("details",
			slog.String("file_name", "usuarios.csv"),
			slog.Int("duration_ms", 250),
			slog.Int("number_of_lines_loaded", 333),
		),
	)
}
```

---

## Child Logging

Com esta funcionalidade, é possível adicionar a vários eventos de log, atributos de um determinado
escopo, sem que seja necessário repetí-los no código.

```go
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
```