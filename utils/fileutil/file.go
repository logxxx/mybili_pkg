package fileutil

import (
	"fmt"
	"github.com/logxxx/mybili_pkg/utils/log"
	"os"
	"path/filepath"
	"strings"
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

func WriteToFileWithRename(dir, fileName string, data []byte) error {
	dir, fileName = getValidPath(dir, fileName)
	return WriteToFile(dir, fileName, data)
}

func getValidPath(dir, fileNameWithExt string) (string, string) {

	if !HasFile(filepath.Join(dir, fileNameWithExt)) {
		return dir, fileNameWithExt
	}

	i := 0
	fileExt := filepath.Ext(fileNameWithExt)
	fileName := strings.TrimRight(fileNameWithExt, fileExt)
	for {
		i++
		fileNameWithExt = fmt.Sprintf("%v_%v%v", fileName, i, fileExt)

		if !HasFile(filepath.Join(dir, fileNameWithExt)) {
			return dir, fileNameWithExt
		}

	}

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

	err = os.MkdirAll(fileDir, 0777)
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

func HasFile(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
