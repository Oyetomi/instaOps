package bot

import (
	"github.com/Oyetomi/instaOps/internal/api"
	"github.com/Oyetomi/instaOps/internal/file"
	"log"
)

const (
	settingsPath = "../cmd/settings/settings.json"
	yamlPath     = "../cmd/settings/config.yaml"
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
	cs := file.ReadConfig(c, yamlPath)
	log.Printf("Logging into %s", cs.Username)
	absPath, err := file.CreateAbsolutePath(settingsPath)
	if err != nil {
		log.Fatalf("Could not create %v", absPath)
	}
	info, err := file.CheckIfFileExists(absPath)
	if err != nil {
		log.Fatalf("%v does not exist", info.Name())
	}
	empty := file.IsEmptyFile(absPath)
	if empty {
		// if file is empty, we do a manual login
		sessionid := api.Login(cs.Username, cs.Password, cs.Verification_code, cs.Proxy, cs.Locale, cs.Timezone)
		// get settings.json and save it into settings.json settings file
		settings := api.GetSettings(sessionid)
		log.Printf("Wrote cookie to %v ", absPath)
		// write cookies to file
		err := file.WriteToFile(absPath, []byte(settings))
		if err != nil {
			log.Fatalf("Could not write cookie to %v", absPath)
		}
	} else {
		// read cookies
		contents, err := file.ReadFileContents(absPath)
		if err != nil {
			log.Fatalf("Could not read cookies from %v", absPath)
		}
		// retrieve sessionid so we can log in
		sessionid = api.SetSettings(string(contents))
		log.Printf("Successfully logged in to %s", cs.Username)
	}
	return sessionid
}
