package fileSystem

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	fsv1 "github.com/vorontsoved/protosFileService/gen/go/fileservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	status2 "google.golang.org/grpc/status"
	"log/slog"
)

type browseElements struct {
	id          int
	filename    string
	created_at  timestamp.Timestamp
	modyfied_at timestamp.Timestamp
}

type FileService interface {
	Upload(
		ctx context.Context,
		file []byte,
		fileName string,
	) (status bool, err error)
	Browse(
		ctx context.Context,
	) (files []browseElements, err error)
}

type serverAPI struct {
	log *slog.Logger
	fsv1.UnimplementedFileServiceServer
	fileService FileService
}

var dict = map[bool]string{
	true:  "Успешная загрузка",
	false: "Ошибка при загрузке файла",
}

func RegisterServerAPI(gRPC *grpc.Server, fileService FileService) {
	fsv1.RegisterFileServiceServer(gRPC, &serverAPI{fileService: fileService})
}

func (s *serverAPI) Upload(ctx context.Context, req *fsv1.FileUploadRequest) (*fsv1.FileUploadResponse, error) {
	file := req.GetFileData()
	fileName := req.GetFileName()
	if len(file) == 0 {
		return &fsv1.FileUploadResponse{Success: false, Message: "Invalid file"}, nil
	}

	status, err := s.fileService.Upload(ctx, file, fileName)
	if err != nil {
		return nil, status2.Error(codes.InvalidArgument, "upload Error")
	}
	return &fsv1.FileUploadResponse{Success: status, Message: dict[status]}, nil
}

func (s *serverAPI) Browse(ctx context.Context, req *fsv1.Empty) (*fsv1.FileBrowseResponse, error) {
	fmt.Println("Browse start")
	files, err := s.fileService.Browse(ctx)
	if err != nil {

	}
	panic("implement me")
}

func (s *serverAPI) Export(ctx context.Context, req *fsv1.FileExportRequest) (*fsv1.FileExportResponse, error) {
	panic("implement me")
}
