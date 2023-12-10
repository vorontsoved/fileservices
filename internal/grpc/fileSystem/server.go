package fileSystem

import (
	"context"
	fsv1 "github.com/vorontsoved/protosFileService/gen/go/fileservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	status2 "google.golang.org/grpc/status"
)

type UploadService interface {
	Upload(
		ctx context.Context,
		file []byte,
		file_name string,
	) (status bool, err error)
}

type serverAPI struct {
	fsv1.UnimplementedFileServiceServer
	uploadService UploadService
}

func RegisterServerAPI(gRPC *grpc.Server, uploadService UploadService) {
	fsv1.RegisterFileServiceServer(gRPC, &serverAPI{uploadService: uploadService})
}

func (s *serverAPI) Upload(ctx context.Context, req *fsv1.FileUploadRequest) (*fsv1.FileUploadResponse, error) {

	file := req.GetFileData()

	if len(file) == 0 {
		return &fsv1.FileUploadResponse{Success: false, Message: "Invalid file"}, nil
	}
	status, err := s.uploadService.Upload(ctx, file, file_name)
	if err != nil {
		return nil, status2.Error(codes.InvalidArgument, "upload Error")
	}
	return &fsv1.FileUploadResponse{Success: status, Message: "Успешно загрузили файл"}, nil
}

func (s *serverAPI) Browse(ctx context.Context, req *fsv1.Empty) (*fsv1.FileBrowseResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Export(ctx context.Context, req *fsv1.FileExportRequest) (*fsv1.FileExportResponse, error) {
	panic("implement me")
}
