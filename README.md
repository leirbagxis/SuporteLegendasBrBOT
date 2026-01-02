# Telegram Bot Modular Exemplar (Go)

Este é um projeto de bot para Telegram altamente modularizado, utilizando a biblioteca `gogram` para interação com a API do Telegram (MTProto) e `gin-gonic` para uma API REST de suporte.

## 🚀 Funcionalidades

- **Modularidade Total:** Código separado por responsabilidades (bot, api, handlers, session).
- **Login por Telefone:** Autenticação via MTProto usando número de telefone e código.
- **API MVP:** Endpoints REST para monitoramento e informações do bot.
- **Extensível:** Fácil de adicionar novos comandos ou rotas de API.

## 📁 Estrutura do Projeto

- `cmd/`: Ponto de entrada da aplicação.
- `internal/bot/`: Lógica central do Telegram.
- `internal/api/`: Servidor web Gin e endpoints.
- `internal/config/`: Gerenciamento de configurações.
- `pkg/`: Código utilitário compartilhado.

## 🛠️ Pré-requisitos

- Go 1.21 ou superior.
- `API_ID` e `API_HASH` obtidos em [my.telegram.org](https://my.telegram.org).

## ⚙️ Configuração

Crie um arquivo `.env` na raiz do projeto ou exporte as variáveis:

```env
TELEGRAM_API_ID=seu_api_id
TELEGRAM_API_HASH=seu_api_hash
TELEGRAM_PHONE=+5511999999999
```

## 🏃 Como Executar

1. Instale as dependências:
   ```bash
   go mod tidy
   ```

2. Execute o projeto:
   ```bash
   go run cmd/main.go
   ```

3. Siga as instruções no terminal para inserir o código de login enviado ao seu Telegram.

## 🌐 Endpoints da API

- `GET /api/v1/status`: Verifica se o sistema está online.
- `GET /api/v1/bot/info`: Retorna informações básicas do bot.

## 📄 Licença

Este projeto é um exemplo para fins educacionais.
