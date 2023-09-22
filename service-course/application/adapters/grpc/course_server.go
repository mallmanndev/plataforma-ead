package grpc

import (
	"context"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseServer struct {
	pb.CoursesServiceServer
	createCourseUseCase *usecases.CreateCourseUseCase
}

func NewCourseServer(db *mongo.Client) *CourseServer {
	peopleRepository := repositories.NewPeopleRepository(db)
	coursesRepo := repositories.NewCourseRepositories(db)
	createCourseUseCase := usecases.NewCreateCourseUseCase(peopleRepository, coursesRepo)
	return &CourseServer{createCourseUseCase: createCourseUseCase}
}

func (cs *CourseServer) Create(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	course, err := cs.createCourseUseCase.Execute(usecases.CreateCourseUseCaseDTO{
		Name:        req.Name,
		Description: req.Description,
		Instructor: usecases.CreateCourseInstructorDTO{
			Id:   req.Instructor.Id,
			Name: req.Instructor.Name,
			Type: req.Instructor.Type,
		},
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	res := &pb.CreateCourseResponse{
		Course: &pb.Course{
			Id:          course.Id(),
			Name:        course.Name(),
			Description: course.Description(),
			Visible:     course.IsVisible(),
			CreatedAt:   course.CreatedAt().String(),
		},
	}

	return res, nil
}
