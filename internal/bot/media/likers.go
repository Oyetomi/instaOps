package media

import (
	"encoding/json"
	"fmt"
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/errors"
	"io"
	"log"
)

type Media struct {
	Pk string `json:"pk"`
}

// GetMediaLikers returns 100 media likers of a particular media
//TODO figure out why it only returns 100 users as opposed to 1000
func GetMediaLikers(sessionid, media_id string, writer io.Writer) {
	var m []Media
	likers, err := api.GetMediaLikers(sessionid, media_id)
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaLikers)
	}
	if err := json.Unmarshal([]byte(likers), &m); err != nil {
		log.Fatal(err)
	}
	for _, v := range m {
		if _, err := fmt.Fprintln(writer, v.Pk); err != nil {
			log.Fatalf("Couldn't write to %v", writer)
		}
	}
}
