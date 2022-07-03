package media

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/errors"
	"log"
)

// UnLikeMedia likes a media
func UnLikeMedia(sessionid, media_id string) {
	_, err := api.UnlikeMedia(sessionid, media_id)
	if err != nil {
		log.Println(errors.ErrMediaCouldNotBeUnliked)
	}
	log.Printf("%v Unliked!", media_id)
}
