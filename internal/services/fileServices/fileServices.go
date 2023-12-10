package fileServices

import (
	"context"
	"log/slog"
)

type FileService struct {
	log     *slog.Logger
	fileSvr FileSaver
}

type FileSaver interface {
	FileSaver(ctx context.Context, filename string) (id int, err error)
}

func New(log *slog.Logger, fileSaver FileSaver) *FileService {
	return &FileService{
		log:     log,
		fileSvr: fileSaver,
	}
}

func (f *FileService) Upload(ctx context.Context, file []byte, filename string) {
	// Вызов метода FileSaver для сохранения файла
	id, err := f.fileSvr.FileSaver(ctx, filename)
	if err != nil {
		f.log.Info("Error saving file: %v\n", err)
		return
	}

	f.log.Info("File saved successfully with ID: %d\n", id)
	return
}
