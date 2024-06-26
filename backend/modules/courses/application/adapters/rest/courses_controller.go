package rest

import (
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/services"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateCourseController struct {
	coursesRepo          ports.CourseRepository
	videosRepo           ports.VideosRepository
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
	videoUploadUseCase   *usecases.VideoUpload
}

func NewCourseServer(db *mongo.Database) *CreateCourseController {
	filesService := services.NewFilesService()
	uuidService := services.NewUUIDService()
	coursesRepo := repositories.NewCourseRepositories(db)
	videosRepo := repositories.NewVideosRepository(db)
	createCourseUseCase := usecases.NewCreateCourseUseCase(coursesRepo)
	updateCourseUseCase := usecases.NewUpdateCourseUseCase(coursesRepo)
	deleteCourseUseCase := usecases.NewDeleteCourseUseCase(coursesRepo)
	createSectionUseCase := usecases.NewCreateSectionUseCase(coursesRepo)
	updateSectionUseCase := usecases.NewUpdateSectionUseCase(coursesRepo)
	deleteSectionUseCase := usecases.NewDeleteSectionUseCase(coursesRepo)
	createItemUseCase := usecases.NewCreateItem(coursesRepo, videosRepo)
	updateItemUseCase := usecases.NewUpdateItem(coursesRepo)
	deleteItemUseCase := usecases.NewDeleteItem(coursesRepo)
	makeCourseVisible := usecases.NewMakeCourseVisible(coursesRepo)
	makeCourseInvisible := usecases.NewMakeCourseInvisible(coursesRepo)
	videoUploadUseCase := usecases.NewVideoUpload(
		filesService,
		uuidService,
		videosRepo,
		apptimer.NewAppTimer(),
	)

	return &CreateCourseController{
		coursesRepo:          coursesRepo,
		videosRepo:           videosRepo,
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
		videoUploadUseCase:   videoUploadUseCase,
	}
}
