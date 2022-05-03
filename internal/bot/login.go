package bot

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/file"
	"github.com/sirupsen/logrus"
)

const (
	settingsPath = "../config/settings.json"
)

func Login(username, password string) (sessionid string) {
	logrus.Println("logging in...")
	// get absolute path to config/settings.json
	settingsPath := file.CreateAbsolutePath(settingsPath)
	file.CheckIfFileExists(settingsPath)
	// check if file is empty
	if file.IsEmptyFile(settingsPath) {
		// if file is empty, we do a manual login
		sessionid = api.Login(username, password)
		// get settings.json and save it into settings.json config file
		settings := api.GetSettings(sessionid)
		logrus.Infof("writing cookie to %v... ", settingsPath)
		file.WriteToFile(settingsPath, []byte(settings))
	} else {
		contents := file.ReadFileContents(settingsPath)
		// retrieve sessionid so we can log in
		sessionid = api.SetSettings(string(contents))
		logrus.Infof("sessionID from %v -> %v", settingsPath, sessionid)
	}
	return sessionid
}
