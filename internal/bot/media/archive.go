package media

import (
	"github.com/Oyetomi/instaOps/internal/api"
)

// ArchiveMedia archives your media
func ArchiveMedia(sessionid, media_id string) {
	api.ArchiveMedia(sessionid, media_id)
}
