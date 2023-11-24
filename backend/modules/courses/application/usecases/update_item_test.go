package usecases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/tests/fixtures"
	"github.com/matheusvmallmann/plataforma-ead/backend/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateItem(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	useCase := usecases.NewUpdateItem(mockCourseRepository)

	t.Run("when_item_is_not_found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindByItemId("item_id").Return(nil, nil)

		course, err := useCase.Execute(usecases.UpdateItemInput{
			Id:          "item_id",
			UserId:      "user_id",
			Title:       "Test item altered",
			Description: "Test item description altered",
		})

		assert.Nil(t, course)
		assert.ErrorContains(t, err, "Item not found.")
	})

	t.Run("when_permission_is_denied", func(t *testing.T) {
		courseFixture := fixtures.NewCourseFixture()

		mockCourseRepository.EXPECT().FindByItemId("item_id").Return(courseFixture, nil)

		course, err := useCase.Execute(usecases.UpdateItemInput{
			Id:          "item_id",
			UserId:      "user_id_invalid",
			Title:       "Test item altered",
			Description: "Test item description altered",
		})

		assert.Nil(t, course)
		assert.ErrorContains(t, err, "Item not found.")
	})

	t.Run("when_update_item_returs_error", func(t *testing.T) {
		courseFixture := fixtures.NewCourseFixture()

		mockCourseRepository.EXPECT().FindByItemId("item_id_1").Return(courseFixture, nil)
		mockCourseRepository.EXPECT().Update(courseFixture).Return(errors.New("Teste."))

		course, err := useCase.Execute(usecases.UpdateItemInput{
			Id:          "item_id_1",
			UserId:      "user_id",
			Title:       "Test item altered",
			Description: "Test item description altered",
		})

		assert.Nil(t, course)
		assert.ErrorContains(t, err, "Teste.")
	})

	t.Run("should_update_item_successfully", func(t *testing.T) {
		courseFixture := fixtures.NewCourseFixture()

		mockCourseRepository.EXPECT().FindByItemId("item_id_1").Return(courseFixture, nil)
		mockCourseRepository.EXPECT().Update(courseFixture).Return(nil)

		course, err := useCase.Execute(usecases.UpdateItemInput{
			Id:          "item_id_1",
			UserId:      "user_id",
			Title:       "Test item altered",
			Description: "Test item description altered",
		})

		item, _ := course.FindItem("item_id_1")

		assert.Nil(t, err)
		if assert.NotNil(t, course) {
			assert.Equal(t, "Test item altered", item.Title())
			assert.Equal(t, "Test item description altered", item.Description())
		}
	})
}
