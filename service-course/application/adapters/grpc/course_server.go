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
	createItemUseCase    *usecases.CreateItem
	updateItemUseCase    *usecases.UpdateItem
	deleteItemUseCase    *usecases.DeleteItem
	makeCourseVisible    *usecases.MakeCourseVisible
	makeCourseInvisible  *usecases.MakeCourseInvisible
}

func NewCourseServer(db *mongo.Database) *CourseServer {
	peopleRepository := repositories.NewPeopleRepository(db)
	coursesRepo := repositories.NewCourseRepositories(db)
	videosRepo := repositories.NewVideosRepository(db)
	createCourseUseCase := usecases.NewCreateCourseUseCase(peopleRepository, coursesRepo)
	updateCourseUseCase := usecases.NewUpdateCourseUseCase(peopleRepository, coursesRepo)
	deleteCourseUseCase := usecases.NewDeleteCourseUseCase(coursesRepo)
	createSectionUseCase := usecases.NewCreateSectionUseCase(coursesRepo)
	updateSectionUseCase := usecases.NewUpdateSectionUseCase(coursesRepo)
	deleteSectionUseCase := usecases.NewDeleteSectionUseCase(coursesRepo)
	createItemUseCase := usecases.NewCreateItem(coursesRepo, videosRepo)
	updateItemUseCase := usecases.NewUpdateItem(coursesRepo)
	deleteItemUseCase := usecases.NewDeleteItem(coursesRepo)
	makeCourseVisible := usecases.NewMakeCourseVisible(coursesRepo)
	makeCourseInvisible := usecases.NewMakeCourseInvisible(coursesRepo)

	return &CourseServer{
		coursesRepo:          coursesRepo,
		createCourseUseCase:  createCourseUseCase,
		updateCourseUseCase:  updateCourseUseCase,
		deleteCourseUseCase:  deleteCourseUseCase,
		createSectionUseCase: createSectionUseCase,
		updateSectionUseCase: updateSectionUseCase,
		deleteSectionUseCase: deleteSectionUseCase,
		createItemUseCase:    createItemUseCase,
		updateItemUseCase:    updateItemUseCase,
		deleteItemUseCase:    deleteItemUseCase,
		makeCourseVisible:    makeCourseVisible,
		makeCourseInvisible:  makeCourseInvisible,
	}
}

func (cs *CourseServer) Create(_ context.Context, req *pb.CreateCourseRequest) (*pb.Course, error) {
	if req.Instructor == nil {
		return nil, status.Error(codes.InvalidArgument, "Instructor is required.")
	}

	course, err := cs.createCourseUseCase.Execute(usecases.CreateCourseUseCaseDTO{
		Name:        req.Name,
		Description: req.Description,
		DiscordUrl:  req.DiscordUrl,
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
		DiscordUrl:  req.DiscordUrl,
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

func (cs *CourseServer) Get(_ context.Context, req *pb.GetCoursesRequest) (*pb.GetCoursesResponse, error) {
	courses, err := cs.coursesRepo.Get(ports.GetCourseFilters{
		Id:      req.Id,
		UserId:  req.UserId,
		Visible: req.Visible,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}
	if courses == nil {
		return &pb.GetCoursesResponse{Courses: nil}, nil
	}

	v := []*pb.Course{}
	for _, value := range courses {
		v = append(v, mappers.CourseEnitiyToGrpc(value))
	}

	response := &pb.GetCoursesResponse{Courses: v}

	return response, nil
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

func (cs *CourseServer) GetSection(_ context.Context, req *pb.GetSectionRequest) (*pb.CourseSection, error) {
	course, err := cs.coursesRepo.FindBySectionId(req.GetId())
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}
	if course == nil {
		return nil, status.Error(codes.NotFound, "Item not found.")
	}

	section := course.FindSection(req.GetId())

	return mappers.SectionEntityToGrpc(section), nil
}

func (cs *CourseServer) CreateItem(_ context.Context, req *pb.CreateItemRequest) (*pb.Course, error) {
	course, err := cs.createItemUseCase.Execute(usecases.CreateItemInput{
		SectionId:   req.SectionId,
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		VideoId:     req.VideoId,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) UpdateItem(_ context.Context, req *pb.UpdateItemRequest) (*pb.Course, error) {
	course, err := cs.updateItemUseCase.Execute(usecases.UpdateItemInput{
		Id:          req.Id,
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) DeleteItem(_ context.Context, req *pb.DeleteItemRequest) (*pb.Course, error) {
	course, err := cs.deleteItemUseCase.Execute(usecases.DeleteItemInput{
		Id:     req.Id,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return mappers.CourseEnitiyToGrpc(course), nil
}

func (cs *CourseServer) GetItem(_ context.Context, req *pb.GetItemRequest) (*pb.CourseItem, error) {
	course, err := cs.coursesRepo.FindByItemId(req.GetId())
	if err != nil {
		return nil, errs.NewGrpcError(err)
	}
	if course == nil {
		return nil, status.Error(codes.NotFound, "Item not found.")
	}

	item, _ := course.FindItem(req.GetId())

	return mappers.ItemEntityToGrpc(item), nil
}

func (cs *CourseServer) MakeVisible(_ context.Context, req *pb.ChangeVisibilityRequest) (*pb.ChangeVisibilityResponse, error) {
	if err := cs.makeCourseVisible.Execute(req.GetId(), req.GetUserId()); err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return &pb.ChangeVisibilityResponse{Ok: true}, nil
}

func (cs *CourseServer) MakeInvisible(_ context.Context, req *pb.ChangeVisibilityRequest) (*pb.ChangeVisibilityResponse, error) {
	if err := cs.makeCourseInvisible.Execute(req.GetId(), req.GetUserId()); err != nil {
		return nil, errs.NewGrpcError(err)
	}

	return &pb.ChangeVisibilityResponse{Ok: true}, nil
}
