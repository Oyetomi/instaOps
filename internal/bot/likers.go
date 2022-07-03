package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Oyetomi/instaOps/internal/api"
	"io"
	"log"
)

type Media struct {
	Pk string `json:"pk"`
}

// GetMediaLikers returns 100 media likers of a particular media
//TODO figure out why it only returns 100 users as opposed to 1000
func GetMediaLikers(media_id, sessionid string, writer io.Writer) {
	var m []Media
	if err := json.Unmarshal([]byte(api.GetMediaLikers(sessionid, media_id)), &m); err != nil {
		log.Fatal(err)
	}
	for _, v := range m {
		if _, err := fmt.Fprintln(writer, v.Pk); err != nil {
			log.Fatalf("Couldn't write to %v", writer)
		}
	}
}
