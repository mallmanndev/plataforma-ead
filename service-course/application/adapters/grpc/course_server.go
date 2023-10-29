package grpc

import (
	"context"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc/mappers"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourseServer struct {
	pb.CoursesServiceServer
	coursesRepo          ports.CourseRepository
	createCourseUseCase  *usecases.CreateCourseUseCase
	updateCourseUseCase  *usecases.UpdateCourseUseCase
	deleteCourseUseCase  *usecases.DeleteCourseUseCase
	createSectionUseCase *usecases.CreateSectionUseCase
	updateSectionUseCase *usecases.UpdateSectionUseCase
	deleteSectionUseCase *usecases.DeleteSectionUseCase
}

func NewCourseServer(db *mongo.Database) *CourseServer {
	peopleRepository := repositories.NewPeopleRepository(db)
	coursesRepo := repositories.NewCourseRepositories(db)
	createCourseUseCase := usecases.NewCreateCourseUseCase(peopleRepository, coursesRepo)
	updateCourseUseCase := usecases.NewUpdateCourseUseCase(peopleRepository, coursesRepo)
	deleteCourseUseCase := usecases.NewDeleteCourseUseCase(coursesRepo)
	createSectionUseCase := usecases.NewCreateSectionUseCase(coursesRepo)
	updateSectionUseCase := usecases.NewUpdateSectionUseCase(coursesRepo)
	deleteSectionUseCase := usecases.NewDeleteSectionUseCase(coursesRepo)
	return &CourseServer{
		coursesRepo:          coursesRepo,
		createCourseUseCase:  createCourseUseCase,
		updateCourseUseCase:  updateCourseUseCase,
		deleteCourseUseCase:  deleteCourseUseCase,
		createSectionUseCase: createSectionUseCase,
		updateSectionUseCase: updateSectionUseCase,
		deleteSectionUseCase: deleteSectionUseCase,
	}
}

func (cs *CourseServer) Create(_ context.Context, req *pb.CreateCourseRequest) (*pb.Course, error) {
	if req.Instructor == nil {
		return nil, status.Error(codes.InvalidArgument, "Instructor is required.")
	}

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

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) Update(_ context.Context, req *pb.UpdateCourseRequest) (*pb.Course, error) {
	if req.Instructor == nil {
		return nil, status.Error(codes.InvalidArgument, "Instructor is required.")
	}

	course, err := cs.updateCourseUseCase.Execute(usecases.UpdateCourseUseCaseDTO{
		Id: req.CourseId,
		Instructor: usecases.UpdateCourseInstructorDTO{
			Id:   req.Instructor.Id,
			Name: req.Instructor.Name,
			Type: req.Instructor.Type,
		},
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) Delete(_ context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	err := cs.deleteCourseUseCase.Execute(usecases.DeleteCourseUseCaseDataDTO{
		Id:     req.CourseId,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	res := &pb.DeleteCourseResponse{
		Message: "Course deleted successfully.",
	}
	return res, nil
}

func (cs *CourseServer) CreateSection(_ context.Context, req *pb.CreateCourseSectionRequest) (*pb.Course, error) {
	course, err := cs.createSectionUseCase.Execute(usecases.CreateSectionDTO{
		CourseId:    req.CourseId,
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) UpdateSection(_ context.Context, req *pb.UpdateCourseSectionRequest) (*pb.Course, error) {
	course, err := cs.updateSectionUseCase.Execute(usecases.UpdateSectionDTO{
		UserId:      req.UserId,
		SectionId:   req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) DeleteSection(_ context.Context, req *pb.DeleteCourseSectionRequest) (*pb.DeleteCourseResponse, error) {
	err := cs.deleteSectionUseCase.Execute(usecases.DeleteSectionDTO{
		UserId:    req.UserId,
		SectionId: req.Id,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	res := &pb.DeleteCourseResponse{
		Message: "Section deleted successfully.",
	}
	return res, nil
}

func (cs *CourseServer) Get(_ context.Context, req *pb.GetCoursesRequest) (*pb.GetCoursesResponse, error) {
	courses, err := cs.coursesRepo.Get(ports.GetCourseFilters{
		Id:           req.Id,
		InstructorId: req.InstructorId,
		Visible:      req.Visible,
	})
	if err != nil {
		return nil, err
	}

	v := []*pb.Course{}
	for _, value := range courses {
		v = append(v, mappers.CourseEnitiyToGrpc(value))
	}

	response := &pb.GetCoursesResponse{Courses: v}

	return response, nil
}
