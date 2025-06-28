package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	Token        = os.Getenv("TOKEN")
	ApiKey       = os.Getenv("API_KEY")
	ApiUrl       = os.Getenv("API_URL")
	Proxy        = os.Getenv("PROXY")
	CoolifyToken = os.Getenv("COOLIFY_TOKEN")
	MongoUrl     = os.Getenv("MONGO_URL")
	DownloadPath = "downloads"
)
