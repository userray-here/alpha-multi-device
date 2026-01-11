package main

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	pref := telebot.Settings{
		Token:  "8474381583:AAFWXHgjEws6NBhvos2JPqHL2nOe21pusJg",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// ===== BUTTON & MENU =====
	menu := &telebot.ReplyMarkup{}

	btnTools := menu.Data("🛠 Tools Menu", "tools")
	btnDevice := menu.Data("📱 Multi Device", "device")
	btnStatus := menu.Data("📊 Status Bot", "status")

	menu.Inline(
		menu.Row(btnTools, btnDevice),
		menu.Row(btnStatus),
	)

	// ===== /START =====
	b.Handle("/start", func(c telebot.Context) error {
		photo := &telebot.Photo{
			File: telebot.FromURL("https://i.imgur.com/your_thumbnail.png"), // ganti thumbnail
			Caption: "*Alpha Assistant*\n\n" +
				"Status : `ACTIVE`\n" +
				"Mode   : Private\n" +
				"Uptime : Running",
		}

		return c.Send(photo, menu, telebot.ModeMarkdown)
	})

	// ===== BUTTON HANDLER =====
	b.Handle(&btnTools, func(c telebot.Context) error {
		return c.Send("Tools menu belum aktif.")
	})

	b.Handle(&btnDevice, func(c telebot.Context) error {
		return c.Send("Multi device belum tersambung.")
	})

	b.Handle(&btnStatus, func(c telebot.Context) error {
		return c.Send("Status bot: ACTIVE\nNo error.")
	})

	log.Println("Bot jalan...")
	b.Start()
}