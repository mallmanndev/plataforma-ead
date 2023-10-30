package grpc

import (
	"context"
	"io"
	"log"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/services"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilesServer struct {
	pb.FileUploadServiceServer
	videosRepository   ports.VideosRepository
	videoUploadUseCase *usecases.VideoUpload
}

func NewFilesServer(db *mongo.Database) *FilesServer {
	filesService := services.NewFilesService()
	uuidService := services.NewUUIDService()
	videosRepository := repositories.NewVideosRepository(db)
	return &FilesServer{
		videosRepository: videosRepository,
		videoUploadUseCase: usecases.NewVideoUpload(
			filesService,
			uuidService,
			videosRepository,
			apptimer.NewAppTimer(),
		),
	}
}

func (s *FilesServer) VideoUpload(stream pb.FileUploadService_VideoUploadServer) error {
	log.Println("Receiving a new video upload request")

	req, err := stream.Recv()
	if err != nil {
		return err
	}

	upload, err := s.videoUploadUseCase.CreateFile(
		req.GetInfo().GetType(),
		req.GetInfo().GetSize(),
		req.GetInfo().GetUserId(),
	)
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

func (s *FilesServer) GetVideo(_ context.Context, req *pb.GetVideoRequest) (*pb.GetVideoResponse, error) {

	video, err := s.videosRepository.Find(req.GetId())
	if err != nil {
		log.Fatal(err)
		return nil, errs.NewGrpcError(err)
	}

	var resolutions []*pb.VideoResolution
	for _, res := range video.GetResolutions() {
		resolutions = append(resolutions, &pb.VideoResolution{
			Resolution: res.Resolution,
			Url:        res.URL,
		})
	}

	response := &pb.GetVideoResponse{
		Id:          video.Id(),
		Type:        video.Type(),
		Status:      video.Status(),
		Size:        video.Size(),
		CreatedAt:   video.CreatedAt().String(),
		UpdatedAt:   video.UpdatedAt().String(),
		Resolutions: resolutions,
		Url:         video.Url(),
	}

	return response, nil
}
