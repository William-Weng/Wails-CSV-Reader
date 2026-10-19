package backend

import (
	"wails-csv-reader/backend/utility"
)

type FileService struct{}

func (service *FileService) ParseFilePath(path string) (rootdir string, filename string, extension string) {
	return utility.ParseFilePath(path)
}

func (service *FileService) ReadCSV(path string) ([][]string, error) {
	return utility.ReadCSV(path)
}

func (service *FileService) ReadResources(path string) (utility.ResourceList, error) {
	return utility.ReadResources(path)
}

func CheckInputFile(path string) error {
	return utility.CheckInputFile(path)
}
