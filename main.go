package main

import (
	"fmt"
	"log"
	"os"
	"songBot/config"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var (
	startTimeStamp = time.Now().Unix()
)

func main() {
	if config.Token == "" || config.ApiKey == "" || config.ApiUrl == "" {
		log.Fatal("Missing environment variables")
	}

	clientConfig := tg.ClientConfig{
		AppID:        6,
		AppHash:      "eb06d4abfb49dc3eeb1aeb98ae0f581e",
		Session:      "session.dat",
		FloodHandler: handleFlood,
	}

	client, err := tg.NewClient(clientConfig)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if _, err = client.Conn(); err != nil {
		log.Fatalf("Failed to connect client: %v", err)
	}

	if err = client.LoginBot(config.Token); err != nil {
		log.Fatalf("Bot login failed: %v", err)
	}

	if err = os.Mkdir("downloads", os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create downloads directory: %v", err)
	}

	initFunc(client)

	me, err := client.GetMe()
	if err != nil {
		log.Fatalf("Failed to get bot information: %v", err)
	}

	uptime := time.Since(time.Unix(startTimeStamp, 0)).String()
	client.Logger.Info(fmt.Sprintf("Authenticated as -> @%s, taken: %s.", me.Username, uptime))
	client.Logger.Info("GoGram version: " + tg.Version)
	client.Idle()
	log.Println("Bot stopped.")
}

func handleFlood(err error) bool {
	wait := tg.GetFloodWait(err)
	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
		return true
	}
	return false
}
