package bot

import (
	"errors"
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Instagram struct {
	Username string
	Password string
}

func (i Instagram) Login() {
	logrus.Println("logging in...")
	// get absolute path to config/settings.json
	if settingsPath, err := filepath.Abs("../config/settings.json"); err != nil {
		logrus.Fatal(err)
	} else {
		// check if settings.json file exists
		if info, err := os.Stat(settingsPath); err == nil {
			logrus.Infof("%v file exists", info.Name())
			logrus.Infof("logging into account with %v", info.Name())
			// check if settings.json file is not empty
			if info.Size() == 0 {
				logrus.Warnf("%v is empty", info.Name())
				// if file is empty, we do a manual login
				sessionid := api.Login(i.Username, i.Password)
				// get settings.json and save it into settings.json config file
				settings := api.GetSettings(sessionid)
				logrus.Infof("writing cookie to %v... ", info.Name())
				if err := os.WriteFile(settingsPath, []byte(settings), 0644); err != nil {
					logrus.Fatalf("could not write config to %v", info.Name())
				}
				logrus.Infof("wrote cookie to %v ", info.Name())
			} else {
				// read contents of settings.json.
				// settings.json , gets sessionid for log in
				if contents, err := os.ReadFile(settingsPath); err != nil {
					logrus.Infof("could not read contents of %v", info.Name())
				} else {
					// retrieve sessionid so we can log in
					sessionid := api.SetSettings(string(contents))
					logrus.Infof("sessionID from %v -> %v", info.Name(), sessionid)
				}
			}
		} else {
			if errors.Is(err, os.ErrNotExist) {
				if _, err := os.Create(settingsPath); err != nil {
					logrus.Fatal("could not create file")
				} else {
					logrus.Infof("settings file created, trying logging in again")
				}
			}
		}
	}
}
