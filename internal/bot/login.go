package bot

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/file"
	log "github.com/sirupsen/logrus"
)

const (
	settingsPath = "../settings/settings.json"
)

func Login(username, password string) (sessionid string) {
	log.Println("logging in...")

	// get absolute path to settings/settings.json
	settingsPath, err := file.CreateAbsolutePath(settingsPath)
	if err != nil {
		log.Errorf("could not create %v", settingsPath)
	}

	info, err := file.CheckIfFileExists(settingsPath)
	if err != nil {
		log.Errorf("%v does not exist", info.Name())
	}

	// check if file is empty
	empty := file.IsEmptyFile(settingsPath)
	if empty {
		// if file is empty, we do a manual login
		sessionid = api.Login(username, password)
		// get settings.json and save it into settings.json settings file
		settings := api.GetSettings(sessionid)
		log.Infof("writing cookie to %v ", settingsPath)
		// write cookies to file
		err := file.WriteToFile(settingsPath, []byte(settings))
		if err != nil {
			log.Fatalf("could not write cookie to %v", settingsPath)
		}
	} else {
		// read cookies
		contents, err := file.ReadFileContents(settingsPath)
		if err != nil {
			log.Fatalf("could not read cookies from %v", settingsPath)
		}
		// retrieve sessionid so we can log in
		sessionid = api.SetSettings(string(contents))
		log.Infof("sessionID from %v -> %v", settingsPath, sessionid)
	}
	return sessionid
}
