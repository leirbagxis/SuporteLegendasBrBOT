# Arquitetura do Projeto Telegram Bot Modular

Este projeto segue uma estrutura modular para garantir escalabilidade e facilidade de manutenção.

## Estrutura de Pastas

```text
/telegram_bot_project
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── bot/              # Lógica do bot Telegram
│   │   ├── handlers/     # Manipuladores de comandos e mensagens
│   │   ├── middleware/   # Middlewares para o bot
│   │   └── session/      # Gerenciamento de sessão e login por telefone
│   ├── api/              # API MVP com Gin-Gonic
│   │   ├── routes/       # Definição de rotas
│   │   ├── controllers/  # Lógica dos endpoints
│   │   └── models/       # Modelos de dados
│   └── config/           # Configurações do sistema
├── pkg/                  # Pacotes compartilhados (utilitários)
├── go.mod                # Dependências do Go
└── README.md             # Documentação do projeto
```

## Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Telegram Lib:** [gogram](https://github.com/amarnathcjd/gogram)
- **Web Framework:** [Gin-Gonic](https://github.com/gin-gonic/gin)
- **Login:** Autenticação por telefone (MTProto)
