package media

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/errors"
	"log"
)

// LikeMedia likes a media
func LikeMedia(sessionid, media_id string) {
	_, err := api.LikeMedia(sessionid, media_id)
	if err != nil {
		log.Println(errors.ErrMediaNotLiked)
	}
	log.Printf("%v liked!", media_id)
}
