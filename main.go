package main

import (
	"fmt"
	"github.com/Oyetomi/instaOps/internal/bot"
)

func main() {
	sesssionID := bot.Login()
	fmt.Println(sesssionID)
}
