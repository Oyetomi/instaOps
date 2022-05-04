package bot

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/file"
	"github.com/sirupsen/logrus"
)

const (
	settingsPath = "./config/settings.json"
)

func Login(username, password string) (sessionid string) {
	logrus.Println("logging in...")
	// get absolute path to config/settings.json
	settingsPath, err := file.CreateAbsolutePath(settingsPath)
	if err != nil {
		logrus.Errorf("could not create %v", settingsPath)
	}
	info, err := file.CheckIfFileExists(settingsPath)
	if err != nil {
		logrus.Errorf("%v does not exist", info.Name())
	}
	// check if file is empty
	empty := file.IsEmptyFile(settingsPath)
	if empty {
		// if file is empty, we do a manual login
		sessionid = api.Login(username, password)
		// get settings.json and save it into settings.json config file
		settings := api.GetSettings(sessionid)
		logrus.Infof("writing cookie to %v ", settingsPath)
		// write cookies to file
		err := file.WriteToFile(settingsPath, []byte(settings))
		if err != nil {
			logrus.Fatalf("could not write cookie to %v", settingsPath)
		}
	} else {
		// read cookies
		contents, err := file.ReadFileContents(settingsPath)
		if err != nil {
			logrus.Fatalf("could not read cookies from %v", settingsPath)
		}
		// retrieve sessionid so we can log in
		sessionid = api.SetSettings(string(contents))
		logrus.Infof("sessionID from %v -> %v", settingsPath, sessionid)
	}
	return sessionid
}
