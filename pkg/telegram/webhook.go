package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
)

const (
	wh     = "https://webhook.site/cb526586-6f6c-4398-87b1-cc313a08d012"
	WHPath = "/lighthouse/telegram/webhook"
)

func (b *Bot) SetWebHookHandler(h func(b *Bot, addr, port string) (error, chan error)) {
	b.wh = h
}

func (b *Bot) WebHook(url, port string) (error, chan error) {
	if b.wh != nil {
		return b.wh(b, url, port)
	}
	return b.StartDefaultWebHook(url, port)
}

func (b *Bot) StartDefaultWebHook(addr, port string) (error, chan error) {
	if port == "" {
		port = "443"
	}

	whErrChan := make(chan error)
	wc, err := tgbotapi.NewWebhook(addr + WHPath + port)
	if err != nil {
		return err, whErrChan
	}

	_, err = b.Api.Request(wc)
	if err != nil {
		return err, whErrChan
	}

	http.HandleFunc(WHPath, func(w http.ResponseWriter, r *http.Request) {
		update, err := b.Api.HandleUpdate(r)
		if err != nil {
			whErrChan <- fmt.Errorf("handle error: %w", err)
		}
		go b.handle(update)
	})

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			whErrChan <- fmt.Errorf("webhook server failed: %w", err)
		}
	}()

	return nil, whErrChan
}

func (b *Bot) StartLocalWebHook(redirect, port string) (error, chan error) {
	whErrChan := make(chan error)

	wc, err := tgbotapi.NewWebhook(redirect)
	if err != nil {
		return err, whErrChan
	}

	_, err = b.Api.Request(wc)
	if err != nil {
		return err, whErrChan
	}

	http.HandleFunc(WHPath, func(w http.ResponseWriter, r *http.Request) {
		update, err := b.Api.HandleUpdate(r)
		if err != nil {
			log.Println("Webhook error:", err)
			return
		}
		go b.handle(update)
	})

	go func() {
		log.Println("webhook listening at " + "http://localhost" + ":" + port + WHPath)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Webhook server failed:", err)
		}
	}()
	return nil, whErrChan
}
