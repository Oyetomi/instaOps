package media

import (
	"github.com/Oyetomi/instaOps/internal/api"
)

// LikeMedia likes a media
func LikeMedia(sessionid, media_id string) {
	api.LikeMedia(sessionid, media_id)
}
