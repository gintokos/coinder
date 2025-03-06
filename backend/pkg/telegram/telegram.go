package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gintokos/coinder/pkg/sl"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	context context.Context
	domain  string
	name    string
	b       *bot.Bot
}

// defines also webhook on route https://domain + webhookpath
func NewBot(ctx context.Context, token, domain, name, webhookpath string) (*Bot, error) {
	mbot := Bot{
		context: ctx,
		domain:  domain,
		name:    name,
	}
	b, err := bot.New(token, bot.WithDefaultHandler(mbot.handler))
	if err != nil {
		return &mbot, err
	}

	// to do gracefull to delete webhook on shutdown
	// b.SetWebhook(ctx, &bot.SetWebhookParams{
	// 	URL: "https://" + domain + webhookpath,
	// })
	// b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{
	// 	DropPendingUpdates: true,
	// })

	commands := []models.BotCommand{
		{Command: "app", Description: "get web app"},
	}

	ok, err := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{Commands: commands})
	if !ok {
		return &mbot, fmt.Errorf("settings commands was failed err: %s", err.Error())
	}

	mbot.b = b
	return &mbot, nil
}

func (mbot *Bot) MustStart() {
	mbot.b.Start(mbot.context)
}

func (mbot *Bot) CreateInvoiceLinkStars(ctx context.Context, amount int) (string, error) {
	return mbot.b.CreateInvoiceLink(ctx, &bot.CreateInvoiceLinkParams{
		Title:       "Donate",
		Description: "Donate to the developer",

		Prices: []models.LabeledPrice{
			{
				Amount: amount,
				Label:  "XTR",
			},
		},

		Payload: fmt.Sprintf("donate_%d", amount),

		Currency: "XTR",
	})
}

func (mbot *Bot) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	ok := mbot.handleCommand(update)
	if ok {
		return
	}
	if update.PreCheckoutQuery != nil {
		slog.Info("pre checkout query")
		b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
			PreCheckoutQueryID: update.PreCheckoutQuery.ID,
			OK:                 true,
			ErrorMessage:       "",
		})
		return
	}

	if update.Message != nil {
		if update.Message.SuccessfulPayment != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      fmt.Sprintf("Payment was successful with payment payload: *%s*", update.Message.SuccessfulPayment.InvoicePayload),
				ParseMode: models.ParseModeMarkdown,
			})
			return
		}
	}

	mbot.HandleWebApp(update)
}

func (mbot *Bot) handleCommand(update *models.Update) bool {
	slog.Info(mbot.domain + " command")
	if update.Message == nil {
		return false
	}

	msg := update.Message.Text
	if !strings.HasPrefix(msg, "/") {
		return false
	}

	command := strings.TrimPrefix(msg, "/")

	var params *bot.SendMessageParams

	switch command {
	case "app":
		params = &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Press btn to open app",
			ReplyMarkup: &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{
							Text: mbot.name,
							WebApp: &models.WebAppInfo{
								URL: "https://" + mbot.domain,
							},
						},
					},
				},
			},
		}
	}

	if params != nil {
		if _, err := mbot.b.SendMessage(mbot.context, params); err != nil {
			slog.Error("failed to send msg in telegram", sl.Err(err))
		}
	}

	return true
}

func (mbot *Bot) HandleWebApp(update *models.Update) bool {
	return true
}
