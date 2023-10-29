package usecases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	fixtures "github.com/matheusvmallmann/plataforma-ead/service-course/tests/fixtures"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	mockVideosRepository := mocks.NewMockVideosRepository(mockCtrl)
	useCase := usecases.NewCreateItem(mockCourseRepository, mockVideosRepository)

	t.Run("when_section_is_not_found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindBySectionId("6dd59058-7014-4bb8-9fb0-2647f82bd028").Return(nil, nil)

		err := useCase.Execute(usecases.CreateItemInput{
			SectionId:   "6dd59058-7014-4bb8-9fb0-2647f82bd028",
			UserId:      "15483ad1-65f7-4027-b092-f9a6a792ac64",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "",
		})

		if assert.NotNil(t, err) {
			assert.ErrorContains(t, err, "Course not found.")
		}
	})

	t.Run("when_permission_is_denied_to_create", func(t *testing.T) {
		course := fixtures.NewCourseFixture()
		mockCourseRepository.EXPECT().FindBySectionId("6dd59058-7014-4bb8-9fb0-2647f82bd028").Return(course, nil)

		err := useCase.Execute(usecases.CreateItemInput{
			SectionId:   "6dd59058-7014-4bb8-9fb0-2647f82bd028",
			UserId:      "15483ad1-65f7-4027-b092-f9a6a792ac64",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "",
		})

		if assert.NotNil(t, err) {
			assert.ErrorContains(t, err, "Permission denied to create item.")
		}
	})

	t.Run("when_video_is_not_found", func(t *testing.T) {
		course := fixtures.NewCourseFixture()
		mockCourseRepository.EXPECT().FindBySectionId("section_id_1").Return(course, nil)
		mockVideosRepository.EXPECT().Find("video_id").Return(nil, nil)

		err := useCase.Execute(usecases.CreateItemInput{
			SectionId:   "section_id_1",
			UserId:      "user_id",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "video_id",
		})

		if assert.NotNil(t, err) {
			assert.ErrorContains(t, err, "Video not found.")
		}
	})

	t.Run("when_video_is_created_by_other_user", func(t *testing.T) {
		course := fixtures.NewCourseFixture()
		video := fixtures.NewUser2Video()
		mockCourseRepository.EXPECT().FindBySectionId("section_id_1").Return(course, nil)
		mockVideosRepository.EXPECT().Find("video_id").Return(video, nil)

		err := useCase.Execute(usecases.CreateItemInput{
			SectionId:   "section_id_1",
			UserId:      "user_id",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "video_id",
		})

		if assert.NotNil(t, err) {
			assert.ErrorContains(t, err, "Permission denied to create item.")
		}
	})

	t.Run("when_create_item_returns_error", func(t *testing.T) {
		course := fixtures.NewCourseFixture()
		video := fixtures.NewVideo()
		mockCourseRepository.EXPECT().FindBySectionId("section_id_1").Return(course, nil)
		mockVideosRepository.EXPECT().Find("video_id").Return(video, nil)
		mockCourseRepository.EXPECT().Update(course).Return(errors.New("Test."))

		err := useCase.Execute(usecases.CreateItemInput{
			SectionId:   "section_id_1",
			UserId:      "user_id",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "video_id",
		})

		if assert.NotNil(t, err) {
			assert.ErrorContains(t, err, "Test.")
		}
	})

	t.Run("should_create_item_successfully", func(t *testing.T) {
		course := fixtures.NewCourseFixture()
		video := fixtures.NewVideo()
		mockCourseRepository.EXPECT().FindBySectionId("section_id_1").Return(course, nil)
		mockVideosRepository.EXPECT().Find("video_id").Return(video, nil)
		mockCourseRepository.EXPECT().Update(course).Return(nil)

		err := useCase.Execute(usecases.CreateItemInput{
			SectionId:   "section_id_1",
			UserId:      "user_id",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "video_id",
		})

		assert.Nil(t, err)
		assert.Equal(t, len(course.Sections()[0].Itens()), 1)
	})
}
