package media

import (
	"github.com/Oyetomi/instaOps/internal/api"
)

// UnLikeMedia unlikes a media
func UnLikeMedia(sessionid, media_id string) {
	api.UnlikeMedia(sessionid, media_id)
}
