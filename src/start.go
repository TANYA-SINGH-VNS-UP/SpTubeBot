package src

import (
	"fmt"
	"time"

	"github.com/amarnathcjd/gogram/telegram"
)

// StartHandle responds to the /start command with a welcome message.
func StartHandle(m *telegram.NewMessage) error {
	bot := m.Client.Me()
	name := m.Sender.FirstName

	response := fmt.Sprintf(
		`👋 Hello <b>%s</b>!

🎧 <b>Welcome to %s</b> — your personal music downloader bot!  
Supports: <b>Spotify</b>, <b>YouTube</b>, <b>Apple Music</b>, and <b>SoundCloud</b>.

🔍 <b>To search:</b> Send a song name or a link.
💬 <b>Inline Search:</b> <code>@%s lofi mood</code>
📥 <b>Group Command:</b> <code>/spotify &lt;url&gt;</code>

Enjoy endless tunes! 🚀`,
		name,
		bot.FirstName,
		bot.Username,
	)

	keyboard := telegram.NewKeyboard().
		AddRow(telegram.Button.URL("💫 Fᴀʟʟᴇɴ Pʀᴏᴊᴇᴄᴛꜱ", "https://t.me/FallenProjects"))

	_, err := m.Reply(response, telegram.SendOptions{
		ReplyMarkup: keyboard.Build(),
	})
	return err
}

// PingHandle responds to the /ping command with the bot's latency.
func PingHandle(m *telegram.NewMessage) error {
	start := time.Now()

	msg, err := m.Reply("⏱️ Pinging...")
	if err != nil {
		return err
	}

	latency := time.Since(start)
	_, err = msg.Edit(fmt.Sprintf("🏓 <b>Pong!</b> <code>%s</code>", latency))
	return err
}
