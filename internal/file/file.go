package file

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
)

type FileInfo fs.FileInfo

func CreateAbsolutePath(FilePath string) (string, error) {
	AbsFilePath, err := filepath.Abs(FilePath)
	if err != nil {
		return AbsFilePath, nil
	}
	return AbsFilePath, err
}

func CreateFile(FilePath string) error {
	_, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfFileExists(FilePath string) (FileInfo, error) {
	info, err := os.Stat(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		CreateFile(FilePath)
	}
	return info, nil
}

func WriteToFile(Filepath string, contents []byte) error {
	if err := os.WriteFile(Filepath, contents, 0644); err != nil {
		return err
	}
	return nil
}

func ReadFileContents(Filepath string) ([]byte, error) {
	contents, err := os.ReadFile(Filepath)
	if err != nil {
		return []byte{}, err
	}
	return []byte(contents), nil
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
