package fileutil

import (
	"pkg.logxxx.com/utils/log"
	"os"
	"path/filepath"
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

func WriteToFile(fileDir, fileName string, data []byte) error {
	file, _, err := GetOrCreateFile(fileDir, fileName)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func GetOrCreateFile(fileDir, fileName string) (*os.File, int64, error) {

	filePath := filepath.Join(fileDir, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		fileInfo, _ := file.Stat()
		return file, fileInfo.Size(), nil
	}

	err = os.MkdirAll(fileDir, 777)
	if err != nil {
		log.Errorf("getOrCreateFile MkdirAll err:%v dir:%v", err, fileDir)
		return nil, 0, err
	}

	file, err = os.Create(filePath)
	if err != nil {
		log.Errorf("getOrCreateFile os.Create err:%v path:%v", err, filePath)
		return nil, 0, err
	}

	return file, 0, nil

}
