package bot

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/errors"
	"github.com/Oyetomi/instaOps/internal/file"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

const (
	yamlFile    = "config.yaml"
	cookiesFile = "cookies.json"
)

type Config struct {
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Verification_code string `yaml:"verification_code"`
	Proxy             string `yaml:"proxy"`
	Locale            string `yaml:"locale"`
	Timezone          string `yaml:"timezone"`
}

func Login(pathToSettings string) (sessionid string) {
	var c Config
	pathToSettings, err := file.CreateAbsolutePath(pathToSettings)
	if err != nil {
		logrus.Print(errors.ErrCouldNotCreateAbsolutePath)
	}
	ok := file.IsEmptyFile(filepath.Join(pathToSettings, yamlFile))
	if ok {
		logrus.Fatal(errors.ErrYamlFileIsEmpty)
	}
	cfg, err := file.ReadYamlConfig(c, filepath.Join(pathToSettings, yamlFile))
	if err != nil {
		logrus.Print(errors.ErrCouldNotReadYamlFile)
	}
	ok = file.IsExistsSettingsFolder(pathToSettings)
	switch {
	case ok:
		logrus.Print("Settings folder found.")
		ok := file.CheckIfFilesExists(filepath.Join(pathToSettings, yamlFile), filepath.Join(pathToSettings, cookiesFile))
		if !ok {
			if err := file.CreateFiles(filepath.Join(pathToSettings, yamlFile), filepath.Join(pathToSettings, cookiesFile)); err != nil {
				logrus.Fatal(errors.ErrCouldNotCreateConfigFiles)
			}
			logrus.Fatalf("Set Up %v at %v", yamlFile, pathToSettings)
		}
		ok = file.IsEmptyFile(filepath.Join(pathToSettings, cookiesFile))
		if !ok {
			logrus.Print("Reading Saved Settings from cookies.json")
			contents, err := file.ReadFileContents(filepath.Join(pathToSettings, cookiesFile))
			if err != nil {
				logrus.Fatal(errors.ErrCouldNotReadCookies)
			}
			sessionid = api.SetSettings(string(contents))
		}
		logrus.Print("Logging in With Credentials...")
		sessionid := api.Login(cfg.Username, cfg.Password, cfg.Verification_code, cfg.Proxy, cfg.Locale, cfg.Timezone)
		settings := api.GetSettings(sessionid)
		if err := file.WriteToFile(filepath.Join(pathToSettings, cookiesFile), []byte(settings)); err != nil {
			logrus.Fatal(errors.ErrCouldNotWriteCookies)
		} else {
			contents, err := file.ReadFileContents(filepath.Join(pathToSettings, cookiesFile))
			if err != nil {
				logrus.Fatal(errors.ErrCouldNotReadCookies)
			}
			sessionid = api.SetSettings(string(contents))
			logrus.Printf("Successfully logged in to %s", cfg.Username)
		}
		return sessionid
	case !ok:
		logrus.Print(errors.ErrSettingsFolderNotFound)
		if err := file.CreateDirectory(pathToSettings); err == nil {
			logrus.Printf("%v created", pathToSettings)
		} else {
			logrus.Fatalf("%v %s", errors.ErrCouldNotCreate, pathToSettings)
		}
	}
	return sessionid
}
