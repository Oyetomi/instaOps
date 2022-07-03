package media

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
func GetMediaLikers(sessionid, media_id string, writer io.Writer) {
	var m []Media
	likers := api.GetMediaLikers(sessionid, media_id)
	if err := json.Unmarshal([]byte(likers), &m); err != nil {
		log.Fatal(err)
	}
	for _, v := range m {
		if _, err := fmt.Fprintln(writer, v.Pk); err != nil {
			log.Fatalf("Couldn't write to %v", writer)
		}
	}
}
