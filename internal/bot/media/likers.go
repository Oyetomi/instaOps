package media

import (
	"encoding/json"
	"fmt"
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/sirupsen/logrus"
	"io"
)

type Media struct {
	Pk string `json:"pk"`
}

// GetMediaLikers return media likers of a particular media
func GetMediaLikers(sessionid, media_id string, writer io.Writer) {
	var m []Media
	likers := api.GetMediaLikers(sessionid, media_id)
	if err := json.Unmarshal([]byte(likers), &m); err != nil {
		logrus.Fatal(err)
	}
	for _, v := range m {
		if _, err := fmt.Fprintln(writer, v.Pk); err != nil {
			logrus.Fatalf("Couldn't write to %v", writer)
		}
	}
	logrus.Println("info written successfully")
}
