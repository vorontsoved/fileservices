package fileServices

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	_ "fileservices/internal/grpc/fileSystem"
)

type FileService struct {
	log     *slog.Logger
	fileSvr FileSaver
}

type FileSaver interface {
	FileSaver(ctx context.Context, fileName string) (id int, err error)
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
	fmt.Printf("Файл сохранен в %s\n", filePath)
	if _, err := f.fileSvr.FileSaver(ctx, filename); err != nil {
		f.log.Error("Ошибка в FileSaver: !err\n", err)
		return false, err
	}
	return true, nil
}

func (f *FileService) Browse(ctx context.Context) {[]browseElements

}
