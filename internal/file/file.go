package file

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func CreateAbsolutePath(FilePath string) string {
	AbsFilePath, err := filepath.Abs(FilePath)
	if err != nil {
		logrus.Fatal(err)
	}
	return AbsFilePath
}

func CreateFile(FilePath string) {
	if _, err := os.Create(FilePath); err != nil {
		logrus.Fatalf("could not create %v", FilePath)
	} else {
		logrus.Infof("%v file created", FilePath)
	}
}

func CheckIfFileExists(FilePath string) {
	if info, err := os.Stat(FilePath); err == nil {
		logrus.Infof("%v file exists", info.Name())
	} else {
		if errors.Is(err, os.ErrNotExist) {
			CreateFile(FilePath)
		}
	}
}

func WriteToFile(Filepath string, contents []byte) {
	if err := os.WriteFile(Filepath, contents, 0644); err != nil {
		logrus.Fatalf("could not write contents to %v", Filepath)
	}
	logrus.Infof("wrote contents to %v ", Filepath)
}

func ReadFileContents(Filepath string) string {
	contents, err := os.ReadFile(Filepath)
	if err != nil {
		logrus.Infof("could not read contents of %v", Filepath)
	}
	return string(contents)
}

func IsEmptyFile(FilePath string) bool {
	if info, err := os.Stat(FilePath); err == nil {
		if info.Size() == 0 {
			logrus.Warnf("%v is empty", FilePath)
		} else {
			return false
		}
	}
	return true
}
