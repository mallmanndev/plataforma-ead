package grpc

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/services"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
)

type FilesServer struct {
	pb.FileUploadServiceServer
	videoUploadUseCase *usecases.VideoUpload
}

func NewFilesServer(db *mongo.Database) *FilesServer {
	filesService := services.NewFilesService()
	uuidService := services.NewUUIDService()
	videosRepository := repositories.NewVideosRepository(db)
	return &FilesServer{
		videoUploadUseCase: usecases.NewVideoUpload(
			filesService,
			uuidService,
			videosRepository,
			apptimer.NewAppTimer(),
		),
	}
}

func (s *FilesServer) VideoUpload(stream pb.FileUploadService_VideoUploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	upload, err := s.videoUploadUseCase.CreateFile(req.GetInfo().GetType(), req.GetInfo().GetSize())
	if err != nil {
		return errs.NewGrpcError(err)
	}

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		chunk := req.GetChunk()
		if err = upload.SendChunk(chunk); err != nil {
			return errs.NewGrpcError(err)
		}
	}

	video, err := upload.Execute()
	if err != nil {
		return errs.NewGrpcError(err)
	}

	err = stream.SendAndClose(
		&pb.VideoUploadResponse{
			Id: video.Id(),
		},
	)
	if err != nil {
		return errs.NewGrpcError(err)
	}

	return nil
}
