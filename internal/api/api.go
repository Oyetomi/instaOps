package api

import (
	"encoding/json"
	"github.com/Oyetomi/instaOps/internal/errors"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
)

var client *resty.Client

type Resp struct {
	Detail string `json:"detail"`
}

var R Resp

func init() {
	client = resty.New()
	client.SetBaseURL("http://localhost:8000")
	client.SetContentLength(true)
}

// GetApiVersion gets current API version .
func GetApiVersion() string {
	resp, err := client.R().Get("/version")
	if err != nil {
		log.Println(errors.ErrCouldNotGetAPIVersion)
	}
	return resp.String()
}

// Login logs into instagram with a valid username and password.
func Login(username, password, verification_code, proxy, locale, timezone string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"username":          username,
			"password":          password,
			"verification_code": verification_code,
			"proxy":             proxy,
			"timezone":          timezone,
		}).Post("/auth/login")
	if err != nil {
		log.Fatal(errors.ErrLoginFailed)
	}
	if resp.StatusCode() != 200 {
		log.Println(resp.String())
	}
	return resp.String()
}

// GetSettings retrieves cookies in json format.
func GetSettings(sessionid string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"sessionid": sessionid,
		}).Get("/auth/settings/get")
	if err != nil {
		log.Fatal(errors.ErrCouldNotGetSettings)
	}
	if resp.StatusCode() != 200 {
		log.Println(resp.String())
	}
	return resp.String()
}

// SetSettings authenticate into instagram with a valid cookie
func SetSettings(settings string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"settings": settings,
		}).Post("/auth/settings/set")
	if err != nil {
		log.Fatal(errors.ErrCouldNotSetSettings)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetTimelineFeed gets instagram feed of current user.
func GetTimelineFeed(sessionid string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"sessionid": sessionid,
		}).Get("/auth/timeline_feed")
	if err != nil {
		log.Println(errors.ErrCouldNotGetFeed)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaID converts media_pk to mediaID
func GetMediaID(media_pk int) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"media_pk": strconv.Itoa(media_pk),
		}).Get("/media/id")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaID)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaPK gets short media id/pk
func GetMediaPK(media_id string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"media_id": media_id,
		}).Get("/media/pk")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaPK)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaPKFromCode return mediaPK from code.
//Example: 45818965 returns "250272944479929"
func GetMediaPKFromCode(code string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"code": code,
		}).Get("/media/pk_from_code")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaPK)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaPKFromURL returns media PK from url
func GetMediaPKFromURL(url string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"url": url,
		}).Get("/media/pk_from_url")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaPK)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaInfoByPK retrieves information about a media in json format by using media pk
func GetMediaInfoByPK(sessionid string, pk int) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"pk":        strconv.Itoa(pk),
		}).Post("/media/info")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaInfoByPk)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetUserMedias returns  specified amount of a users media information
func GetUserMedias(sessionid string, user_id, amount int) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   strconv.Itoa(user_id),
			"amount":    strconv.Itoa(amount),
		}).Post("/media/user_medias")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUserMedias)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// DeleteMediaByMediaID delete a media by MediaID
func DeleteMediaByMediaID(sessionid, media_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
		}).Post("/media/delete")
	if err != nil {
		log.Println(errors.ErrCouldNotDeleteMedia)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// EditMedia edits a media
func EditMedia(sessionid, media_id, caption, title string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
			"caption":   caption,
			"title":     title,
		}).Post("/media/edit")
	if err != nil {
		log.Println(errors.ErrCouldNotEditMedia)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaAuthor returns info about the author of the media
func GetMediaAuthor(sessionid, media_pk string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_pk":  media_pk,
		}).Post("/media/user")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaAuthor)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetMediaOembed Return info about media and user from post URL
func GetMediaOembed(sessionid, url string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"url":       url,
		}).Post("/media/oembed")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaInfo)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// LikeMedia like a media
func LikeMedia(sessionid, media_id string) {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
		}).Post("/media/like")
	if resp.StatusCode() != 200 {
		err = json.Unmarshal([]byte(resp.String()), &R)
		if err != nil {
			log.Println(err)
		}
		log.Println(R.Detail)
	} else {
		log.Printf("media %v liked!", media_id)
	}
}

// UnlikeMedia unlikes a media
func UnlikeMedia(sessionid, media_id string) {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
		}).Post("/media/unlike")
	if resp.StatusCode() != 200 {
		err = json.Unmarshal([]byte(resp.String()), &R)
		if err != nil {
			log.Println(err)
		}
		log.Println(R.Detail)
	} else {
		log.Printf("media %v un-liked!", media_id)
	}
}

//TODO: implement MarkMediaAsSeen

// GetMediaLikers gets list of users who liked a media
func GetMediaLikers(sessionid, media_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
		}).Post("/media/likers")
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// ArchiveMedia archives a media
func ArchiveMedia(sessionid, media_id string) {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
		}).Post("/media/archive")
	if resp.StatusCode() != 200 {
		err = json.Unmarshal([]byte(resp.String()), &R)
		if err != nil {
			log.Println(err)
		}
		log.Println(R.Detail)
	} else {
		log.Printf("media %v successfully archived", media_id)
	}
}

// UnArchiveMedia unarchives a media
func UnArchiveMedia(sessionid, media_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"media_id":  media_id,
		}).Post("/media/unarchive")
	if err != nil {
		log.Println(errors.ErrCouldNotUnArchiveMedia)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UploadVideoToStory uploads video to instagram story
func UploadVideoToStory(sessionid, file string) string {
	resp, err := client.R().SetFiles(
		map[string]string{
			"file": file,
		}).SetFormData(
		map[string]string{
			"sessionid": sessionid,
		}).Post("/video/upload_to_story")
	if err != nil {
		log.Println(errors.ErrCouldNotUploadToStory)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UploadVideoToStoryByURL uploads a video specified by URL to instagram story.
// tested with .AVI format https://media-upload.net/uploads/5tfYymMulc9q.avi
func UploadVideoToStoryByURL(sessionid, url string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"url":       url,
		}).Post("/video/upload_to_story/by_url")
	if err != nil {
		log.Println(errors.ErrCouldNotUploadToStory)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// DownloadVideo downloads a video from instagram by specified media_pk
// folder takes a folder path to save the video.
// set returnFile to true to save the video locally
func DownloadVideo(sessionid, media_pk, folder, returnFile string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid":  sessionid,
			"media_pk":   media_pk,
			"folder":     folder,
			"returnFile": returnFile,
		}).Post("/video/download")
	if err != nil {
		log.Println(errors.ErrCouldNotDownloadVideo)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// DownloadVideoByURL downloads a video from instagram by specified URL
// folder takes a folder path to save the video.
// set returnFile to true to save the video locally
//TODO: fix index out of range bug
func DownloadVideoByURL(sessionid, url, filename, folder string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"url":       url,
			"filename":  filename,
			"folder":    folder,
		}).Post("/video/download/by_url")
	if err != nil {
		log.Println(errors.ErrCouldNotDownloadVideo)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UploadPhoto uploads photo to instagram
func UploadPhoto(sessionid, file, caption string) string {
	resp, err := client.R().SetFiles(map[string]string{
		"file": file,
	}).SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"caption":   caption,
		}).Post("/photo/upload")
	if err != nil {
		log.Println(errors.ErrCouldNotUploadPhoto)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UploadPhotoByURL uploads photo to instagram
// image should be viewable from your browser
// image URL ending with .jpg works just fine
func UploadPhotoByURL(sessionid, url, caption string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"url":       url,
			"caption":   caption,
		}).Post("/photo/upload/by_url")
	if err != nil {
		log.Println(errors.ErrCouldNotUploadPhoto)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetUserFollowers returns a specified amount of followers.
func GetUserFollowers(sessionid, user_id, amount string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
			"amount":    amount,
		}).Post("/user/followers")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUserFollowers)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetUserFollowing returns a specified amount of following.
func GetUserFollowing(sessionid, user_id, amount string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
			"amount":    amount,
		}).Post("/user/following")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUsersFollowing)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetUserInfo gets user information
func GetUserInfo(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/info")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUserInfo)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetUserInfoByUsername gets user information by username
func GetUserInfoByUsername(sessionid, username string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"username":  username,
		}).Post("/user/info_by_username")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUserInfo)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// FollowUser follows a user via user_id
func FollowUser(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/follow")
	if err != nil {
		log.Println(errors.ErrCouldNotFollowUser)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UnFollowUser unfollows a user
func UnFollowUser(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/unfollow")
	if err != nil {
		log.Println(errors.ErrCouldNotUnFollowUser)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetIDFromUsername return userID from username
func GetIDFromUsername(sessionid, username string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"username":  username,
		}).Post("/user/id_from_username")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUserIdFromUsername)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// GetUsernameFromID returns username from user_id
func GetUsernameFromID(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/username_from_id")
	if err != nil {
		log.Println(errors.ErrCouldNotGetUsernameFromUserId)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// RemoveFollower removes a follower from your list of followers
func RemoveFollower(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/remove_follower")
	if err != nil {
		log.Println(errors.ErrCouldNotRemoveFollower)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// MutePostsOfFollower mutes a posts of a follower
func MutePostsOfFollower(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/mute_posts_from_follow")
	if err != nil {
		log.Println(errors.ErrCouldNotMutePosts)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UnMutePostsOfFollower unmutes a posts of a follower
func UnMutePostsOfFollower(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/unmute_posts_from_follow")
	if err != nil {
		log.Println(errors.ErrCouldNotUnmutePosts)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// MuteStoriesOfAFollower mutes a follower's stories
func MuteStoriesOfAFollower(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/mute_stories_from_follow")
	if err != nil {
		log.Println(errors.ErrCouldNotMuteStories)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// UnMuteStoriesOfAFollower unmutes a follower's stories
func UnMuteStoriesOfAFollower(sessionid, user_id string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid": sessionid,
			"user_id":   user_id,
		}).Post("/user/unmute_stories_from_follow")
	if err != nil {
		log.Println(errors.ErrCouldNotUnMuteStories)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// DownloadIgtv downloads igtv video
func DownloadIgtv(sessionid, media_pk, folder, returnFile string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid":  sessionid,
			"media_pk":   media_pk,
			"folder":     folder,
			"returnFile": returnFile,
		}).Post("/igtv/download")
	if err != nil {
		log.Println(errors.ErrCouldNotDownloadIgtv)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

// DownloadIgtvByURL downloads igtv video by URL
// TODO: fix error
func DownloadIgtvByURL(sessionid, url, filename, folder, returnFile string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"sessionid":  sessionid,
			"url":        url,
			"filename":   filename,
			"folder":     folder,
			"returnFile": returnFile,
		}).Post("/igtv/download/by_url")
	if err != nil {
		log.Println(errors.ErrCouldNotDownloadIgtv)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}
