package src

import (
	"github.com/amarnathcjd/gogram/telegram"
)

// GoGramVersion responds with the current GoGram version.
func GoGramVersion(m *telegram.NewMessage) error {
	_, err := m.Reply("🤖 <b>GoGram Version:</b> " + telegram.Version)
	return err
}
