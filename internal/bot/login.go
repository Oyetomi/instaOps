package bot

import (
	"errors"
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	settingsPath = "../config/settings.json"
)

func createAbsolutePath(FilePath string) string {
	AbsFilePath, err := filepath.Abs(FilePath)
	if err != nil {
		logrus.Fatal(err)
	}
	return AbsFilePath
}

func createFile(FilePath string) {
	if _, err := os.Create(FilePath); err != nil {
		logrus.Fatalf("could not create %v", FilePath)
	} else {
		logrus.Infof("%v file created", FilePath)
	}
}

func checkIfFileExists(FilePath string) {
	if info, err := os.Stat(FilePath); err == nil {
		logrus.Infof("%v file exists", info.Name())
	} else {
		if errors.Is(err, os.ErrNotExist) {
			createFile(FilePath)
		}
	}
}

func writeToFile(Filepath string, contents []byte) {
	if err := os.WriteFile(Filepath, contents, 0644); err != nil {
		logrus.Fatalf("could not write contents to %v", Filepath)
	}
	logrus.Infof("wrote contents to %v ", Filepath)
}

func readFileContents(Filepath string) string {
	contents, err := os.ReadFile(Filepath)
	if err != nil {
		logrus.Infof("could not read contents of %v", Filepath)
	}
	return string(contents)
}

func isEmptyFile(FilePath string) bool {
	if info, err := os.Stat(FilePath); err == nil {
		if info.Size() == 0 {
			logrus.Warnf("%v is empty", FilePath)
		} else {
			return false
		}
	}
	return true
}

func Login(username, password string) (sessionid string) {
	logrus.Println("logging in...")
	// get absolute path to config/settings.json
	settingsPath := createAbsolutePath(settingsPath)
	checkIfFileExists(settingsPath)
	// check if file is empty
	if isEmptyFile(settingsPath) {
		// if file is empty, we do a manual login
		sessionid = api.Login(username, password)
		// get settings.json and save it into settings.json config file
		settings := api.GetSettings(sessionid)
		logrus.Infof("writing cookie to %v... ", settingsPath)
		writeToFile(settingsPath, []byte(settings))
	} else {
		contents := readFileContents(settingsPath)
		// retrieve sessionid so we can log in
		sessionid = api.SetSettings(string(contents))
		logrus.Infof("sessionID from %v -> %v", settingsPath, sessionid)
	}
	return sessionid
}
