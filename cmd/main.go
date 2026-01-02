package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"telegram_bot_project/internal/api/routes"
	"telegram_bot_project/internal/bot/handlers"
	"telegram_bot_project/internal/bot/session"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente (opcional, mas recomendado)
	_ = godotenv.Load()

	// Configurações (Pode ser movido para internal/config)
	apiIDStr := os.Getenv("TELEGRAM_API_ID")
	apiHash := os.Getenv("TELEGRAM_API_HASH")
	phoneNumber := os.Getenv("TELEGRAM_PHONE")

	if apiIDStr == "" || apiHash == "" || phoneNumber == "" {
		log.Fatal("TELEGRAM_API_ID, TELEGRAM_API_HASH e TELEGRAM_PHONE devem estar configurados no ambiente")
	}

	apiID, _ := strconv.ParseInt(apiIDStr, 10, 32)

	// 1. Inicializar Cliente Telegram
	client, err := session.NewClient(int32(apiID), apiHash)
	if err != nil {
		log.Fatalf("Erro ao criar cliente: %v", err)
	}

	// 2. Login por Telefone
	err = session.Login(client, phoneNumber)
	if err != nil {
		log.Fatalf("Erro no login: %v", err)
	}

	// 3. Registrar Handlers
	handlers.RegisterHandlers(client)

	// 4. Iniciar API em uma goroutine
	go func() {
		fmt.Println("Iniciando API Gin na porta 8080...")
		r := routes.SetupRouter(client)
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Erro ao rodar API: %v", err)
		}
	}()

	// 5. Rodar o Bot
	fmt.Println("Bot Telegram está rodando...")
	client.Idle()
}
