package file

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateAbsolutePath(FilePath string) (string, error) {
	AbsFilePath, err := filepath.Abs(FilePath)
	if err != nil {
		return AbsFilePath, err
	}
	return AbsFilePath, nil
}

func CreateFile(FilePath string) error {
	_, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	return nil
}

func CreateFiles(FilePath ...string) error {
	for _, f := range FilePath {
		_, err := os.Create(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckIfFileExists(FilePath string) bool {
	_, err := os.Stat(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func CheckIfFilesExists(FilePath ...string) bool {
	for _, f := range FilePath {
		_, err := os.Stat(f)
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
	}
	return true
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
			return true
		}
	}
	return false
}

func ReadYamlConfig[T any](config T, yamlFile string) (*T, error) {
	bytesOut, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return &config, err
	}
	if err := yaml.Unmarshal(bytesOut, &config); err != nil {
		return &config, err
	}
	return &config, nil
}

func IsExistsSettingsFolder(path string) bool {
	ok := CheckIfFileExists(path)
	if ok {
		return true
	}
	return false
}

func CreateDirectory(path string) error {
	if err := os.Mkdir(path, 0777); err != nil {
		return err
	}
	return nil
}
