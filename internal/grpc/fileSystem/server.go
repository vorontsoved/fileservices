package fileSystem

import (
	"context"
	fsv1 "github.com/vorontsoved/protosFileService/gen/go/fileservices"
	"google.golang.org/grpc"
)

type serverAPI struct {
	fsv1.UnimplementedFileServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	fsv1.RegisterFileServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Upload(ctx context.Context, req *fsv1.FileUploadRequest) (*fsv1.FileUploadResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Browse(ctx context.Context, req *fsv1.Empty) (*fsv1.FileBrowseResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Export(ctx context.Context, req *fsv1.FileExportRequest) (*fsv1.FileExportResponse, error) {
	panic("implement me")
}
