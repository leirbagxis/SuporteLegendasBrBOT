package session

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"
)

func NewClient(apiID int32, apiHash string) (*telegram.Client, error) {
	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:   apiID,
		AppHash: apiHash,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Login(client *telegram.Client, phoneNumber string) error {
	fmt.Printf("Iniciando login para o número: %s\n", phoneNumber)

	_, err := client.Login(phoneNumber)
	if err != nil {
		return fmt.Errorf("erro ao enviar código: %v", err)
	}

	client.SendMessage("me", "Login realizado com sucesso!")
	fmt.Println("Login realizado com sucesso!")
	return nil
}
