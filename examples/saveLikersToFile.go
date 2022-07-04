package examples

import (
	"github.com/Oyetomi/instaOps/internal/bot/media"
	"log"
	"os"
)

func SaveLikersToFile(sessionid, media_id string, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("%v could not create file", file.Name())
	}
	media.GetMediaLikers(sessionid, media_id, file)
}
