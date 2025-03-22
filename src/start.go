package src

import (
	"fmt"

	"time"

	"github.com/amarnathcjd/gogram/telegram"
)

func StartHandle(m *telegram.NewMessage) error {
	me := m.Client.Me()
	text := fmt.Sprintf(
		"👋 Hey %s!\n\n"+
			"🎵 <b>Welcome to the Ultimate Music Downloader Bot!</b>\n"+
			"Download high-quality songs instantly. Just follow these simple steps:\n\n"+
			"🔹 <b>Send me the name of any song</b>, and I'll fetch it for you.\n"+
			"🔹 You can also use <b>inline mode</b> to search for songs:\n"+
			"Example: <code>@%s mood lofi</code>\n\n"+
			"🔹 <b>Send me the Spotify URL of any song</b>, and I'll download it for you.\n"+
			"🔹 /spotify url: if you want to download a song from Spotify in Group.\n"+
			"✨ Enjoy your music! Happy listening! 🎶",
		m.Sender.FirstName, me.Username,
	)

	opts := telegram.SendOptions{
		ReplyMarkup: telegram.NewKeyboard().AddRow(
			telegram.Button.URL("Fᴀʟʟᴇɴ PʀᴏJᴇᴄᴛs", "https://t.me/FallenProjects"),
		).Build(),
	}
	_, _ = m.Reply(text, opts)
	return nil
}

func PingHandle(m *telegram.NewMessage) error {
	startTime := time.Now()
	sentMessage, _ := m.Reply("Pinging...")
	fmt.Println("Pong!")
	_, err := sentMessage.Edit(fmt.Sprintf("<code>Pong!</code> <code>%s</code>", time.Since(startTime).String()))
	return err
}
