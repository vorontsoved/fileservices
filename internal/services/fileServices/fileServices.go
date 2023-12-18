package fileServices

import (
	"context"
	"fileservices/internal/grpc/fileSystem"
	"fmt"
	"log/slog"
	"os"
)

type FileService struct {
	log     *slog.Logger
	fileSvr FileSaver
}

type FileSaver interface {
	FileSaver(ctx context.Context, fileName string) (id int, err error)
	GetNameFiles(ctx context.Context) (files []fileSystem.BrowseElements, err error)
	GetFile(ctx context.Context, fileId int64) (file []byte, err error)
}

func New(log *slog.Logger, fileSaver FileSaver) *FileService {
	return &FileService{
		log:     log,
		fileSvr: fileSaver,
	}
}

func (f *FileService) Upload(ctx context.Context, file []byte, filename string) (status bool, err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		f.log.Error("Ошибка в получении текущей директории")
	}
	filePath := fmt.Sprintf("%s/SaveFiles/%s", currentDir, filename)

	if _, err := os.Stat(filePath); err == nil {
		f.log.Error("Файл с таким именем уже существует")
		return false, err
	}

	fileHandle, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return false, err
	}
	defer fileHandle.Close()

	_, err = fileHandle.Write(file)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return false, err
	}
	if _, err := f.fileSvr.FileSaver(ctx, filename); err != nil {
		f.log.Error("Ошибка в FileSaver: !err\n", err)
		return false, err
	}
	return true, nil
}

func (f *FileService) Browse(ctx context.Context) (files []fileSystem.BrowseElements, err error) {
	resFiles, err := f.fileSvr.GetNameFiles(ctx)
	if err != nil {
		return []fileSystem.BrowseElements{}, err
	}
	return resFiles, nil
}

func (f *FileService) Export(ctx context.Context, fileId int64) (file []byte, err error) {
	readFile, err := f.fileSvr.GetFile(ctx, fileId)
	if err != nil {
		return readFile, err
	}
	return readFile, nil
}
