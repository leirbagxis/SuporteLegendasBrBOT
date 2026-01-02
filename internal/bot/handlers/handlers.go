package handlers

import (
	"github.com/amarnathcjd/gogram/telegram"
)

func RegisterHandlers(client *telegram.Client) {
	client.On("message", func(m *telegram.NewMessage) error {
		if m.Text() == "/start" {
			_, err := m.Reply("Olá! Eu sou um bot modular exemplar feito em Go.")
			return err
		}
		if m.Text() == "/ping" {
			_, err := m.Reply("Pong! 🏓")
			return err
		}
		return nil
	})
}
