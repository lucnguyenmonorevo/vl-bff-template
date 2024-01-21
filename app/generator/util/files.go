package util

import (
	"os"
)

func Mkdirs(path string) error {
	fileInfo, err := os.Lstat("./")
	if err != nil {
		return err
	}

	fileMode := fileInfo.Mode()
	unixPerms := fileMode & os.ModePerm
	if err := os.MkdirAll(path, unixPerms); err != nil {
		return err
	}

	return nil
}

func ExistFile(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func ReadFile(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	text := string(bytes)
	return text, nil
}

func WriteFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return err
	}
	return nil
}

func AppendFile(filePath string, data []byte) error {
	exist := ExistFile(filePath)

	if exist {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}

		defer file.Close()

		if _, err = file.Write(data); err != nil {
			return err
		}
	} else {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.Write(data); err != nil {
			return err
		}
	}

	return nil
}
