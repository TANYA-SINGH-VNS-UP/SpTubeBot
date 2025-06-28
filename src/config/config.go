package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Tokens       []string
	ApiKey       = os.Getenv("API_KEY")
	ApiUrl       = os.Getenv("API_URL")
	Proxy        = os.Getenv("PROXY")
	CoolifyToken = os.Getenv("COOLIFY_TOKEN")
	DownloadPath = "downloads"
)

func init() {
	tokensEnv := os.Getenv("TOKENS")
	if tokensEnv != "" {
		for _, t := range strings.Split(tokensEnv, ",") {
			t = strings.TrimSpace(t)
			if t != "" {
				Tokens = append(Tokens, t)
			}
		}
	}
}
