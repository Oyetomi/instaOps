package main

import (
	"github.com/Oyetomi/instaOps/examples"
	"github.com/Oyetomi/instaOps/internal/bot"
)

const (
	pathToSettings = "settings"
)

func main() {
	sessionid := bot.Login(pathToSettings)
	examples.SaveLikersToFile(sessionid, "2526987517325974533_13579001845", "../output/likers.txt")
}
