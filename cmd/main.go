package main

import (
	"fmt"
	"github.com/Oyetomi/instaOps/internal/api"
)

func main() {
	mediaID := api.GetUserFollowing("45843020069%3AK28p6piMElUFXk%3A18", "1277259526", "2")
	fmt.Println(mediaID)
}
