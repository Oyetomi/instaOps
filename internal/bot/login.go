package bot

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/file"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

const (
	settingsPath      = "../cmd/settings"
	yamlFile          = "config.yaml"
	settingsFile      = "settings.json"
	sessionIDFilePath = "./sessionID.txt"
)

type Config struct {
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Verification_code string `yaml:"verification_code"`
	Proxy             string `yaml:"proxy"`
	Locale            string `yaml:"locale"`
	Timezone          string `yaml:"timezone"`
	SettingsPath      string `yaml:"settingsPath"`
	YamlPath          string `yaml:"yamlPath"`
}

func Login() (sessionid string) {
	var c Config
	settingsPath, err := file.CreateAbsolutePath(settingsPath)
	if err != nil {
		logrus.Print("Could not create absolute path")
	}
	ok := file.IsEmptyFile(filepath.Join(settingsPath, yamlFile))
	if ok {
		logrus.Fatal("yaml file is empty")
	}
	cfg, err := file.ReadYamlConfig(c, filepath.Join(settingsPath, yamlFile))
	if err != nil {
		logrus.Print("Could not read yaml config")
	}
	ok = file.IsExistsSettingsFolder(settingsPath)
	switch {
	case ok:
		logrus.Print("Settings folder found.")
		ok := file.CheckIfFilesExists(filepath.Join(settingsPath, yamlFile), filepath.Join(settingsPath, settingsFile))
		if !ok {
			if err := file.CreateFiles(filepath.Join(settingsPath, yamlFile), filepath.Join(settingsPath, settingsFile)); err != nil {
				logrus.Fatal("Could not configuration files")
			}
			logrus.Fatalf("Set Up %v at %v", yamlFile, settingsPath)
		}
		ok = file.IsEmptyFile(filepath.Join(settingsPath, settingsFile))
		if !ok {
			logrus.Print("Reading Saved Settings from settings.json")
			contents, err := file.ReadFileContents(filepath.Join(settingsPath, settingsFile))
			if err != nil {
				logrus.Fatalf("Could not read cookies")
			}
			sessionid = api.SetSettings(string(contents))
		}
		logrus.Print("Logging in With Credentials...")
		sessionid := api.Login(cfg.Username, cfg.Password, cfg.Verification_code, cfg.Proxy, cfg.Locale, cfg.Timezone)
		settings := api.GetSettings(sessionid)
		if err := file.WriteToFile(filepath.Join(settingsPath, settingsFile), []byte(settings)); err != nil {
			logrus.Fatal("Could not write cookies")
		} else {
			contents, err := file.ReadFileContents(filepath.Join(settingsPath, settingsFile))
			if err != nil {
				logrus.Fatal("could not read cookies")
			}
			sessionid = api.SetSettings(string(contents))
			logrus.Printf("Successfully logged in to %s", cfg.Username)
		}
	case !ok:
		logrus.Print("Settings folder not found.")
		if err := file.CreateDirectory(settingsPath); err == nil {
			logrus.Printf("%v created", settingsPath)
		} else {
			logrus.Fatalf("%v could not create", settingsPath)
		}
	}
	return sessionid
}
