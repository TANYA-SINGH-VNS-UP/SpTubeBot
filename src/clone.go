package src

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"
	"songBot/src/config"
)

var tokenRegex = regexp.MustCompile(`\b\d{6,}:[\w-]{30,}\b`)

// extractBotToken extracts a Telegram bot token using regex
func extractBotToken(text string) string {
	return tokenRegex.FindString(text)
}

// extractBotIDFromToken gets the numeric ID part from the token
func extractBotIDFromToken(token string) string {
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

func cloneHandle(m *telegram.NewMessage) error {
	const MainBotID int64 = 7805563183
	if m.Client.Me().ID != MainBotID {
		_, _ = m.Reply("🚫 Please send your bot token to the main bot: @SpTubeBot")
		return nil
	}

	token := extractBotToken(m.Text())
	if token == "" {
		_, _ = m.Reply("⚠️ Invalid or missing bot token. Please send a valid token from @BotFather.")
		return nil
	}

	botID := extractBotIDFromToken(token)
	if botID == "" {
		_, _ = m.Reply("❌ Failed to parse Bot ID from token.")
		return nil
	}

	msg, err := m.Reply("🔄 Cloning your bot...")
	if err != nil {
		return err
	}

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:         8,
		AppHash:       "7245de8e747a0d6fbe11f7cc14fcc0bb",
		MemorySession: true,
		SessionName:   fmt.Sprintf("bot_%s", botID),
	})

	if err != nil {
		_, _ = msg.Edit(fmt.Sprintf("❌ Failed to create client: %v", err))
		return nil
	}

	if _, err = client.Conn(); err != nil {
		_, _ = msg.Edit(fmt.Sprintf("❌ Connection error: %v", err))
		return nil
	}

	if err = client.LoginBot(token); err != nil {
		_, _ = msg.Edit(fmt.Sprintf("❌ Bot login failed: %v", err))
		return nil
	}

	if err = config.AddBotToken(m.SenderID(), token); err != nil {
		log.Printf("❌ Failed to store token for user %d: %v", m.SenderID(), err)
		_, _ = msg.Edit("✅ Bot cloned, but failed to save token in DB.\n\n" +
			"" + err.Error())
	} else {
		me := client.Me()
		m.Client.Logger.Info("Bot cloned successfully", "bot_name", me.FirstName, "bot_username", me.Username, "bot_id", me.ID)
		_, _ = msg.Edit(fmt.Sprintf(
			`✅ <b>Your bot was cloned successfully and saved!</b>

<b>🤖 Bot Name:</b> %s
<b>🔗 Username:</b> @%s
<b>🆔 Bot ID:</b> %d`,
			me.FirstName, me.Username, me.ID,
		), telegram.SendOptions{ParseMode: "HTML"})
	}

	InitFunc(client)
	return nil
}
