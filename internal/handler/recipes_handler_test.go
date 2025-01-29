package handler_test

import (
	"testing"
	"time"

	"gororoba/internal/handler"
	testdata_fixtures "gororoba/internal/testdata/fixtures"
	testdata_mocks "gororoba/internal/testdata/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testSetup struct {
	recipeHandler        *handler.RecipesHandler
	recipeRepositoryMock *testdata_mocks.MockRecipeRepositoryInterface
	recipeSuggestionMock *testdata_mocks.MockSuggestionHandlerInterface
}

func setup(t *testing.T) testSetup {
	ctrl := gomock.NewController(t)
	rr := testdata_mocks.NewMockRecipeRepositoryInterface(ctrl)
	sh := testdata_mocks.NewMockSuggestionHandlerInterface(ctrl)
	h := handler.NewRecipesHandler(rr, sh)

	return testSetup{
		recipeHandler:        &h,
		recipeRepositoryMock: rr,
		recipeSuggestionMock: sh,
	}
}

func TestGetRecipesByCategory(t *testing.T) {

	// Given
	s := setup(t)
	c := "Dessert"
	r := testdata_fixtures.GetRecipesWithCategory(c)
	s.recipeRepositoryMock.EXPECT().GetRecipesByCategory(c).Return(r)

	// When
	result := s.recipeHandler.GetRecipesByCategory(c)

	// Then
	assert.GreaterOrEqual(t, len(result), 1)

}

func TestCreateRecipe(t *testing.T) {
	// Given
	s := setup(t)
	r := testdata_fixtures.GetRecipesWithCategory("salad")
	s.recipeRepositoryMock.EXPECT().CreateRecipe(gomock.Any()).Return(nil)

	// When
	result := s.recipeHandler.CreateRecipe(&r[0])

	// Then
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Id)
	assert.NotEmpty(t, result.CreatedAt)
	assert.NotEmpty(t, result.UpdatedAt)
}

func TestGivenACategorySuggestionShouldReturnARecipe(t *testing.T) {
	// Given
	s := setup(t)
	now := time.Now()
	suggestedCategory := "Dessert"
	s.recipeSuggestionMock.EXPECT().GetSuggestedCategoryByTime(now).Return(suggestedCategory)
	s.recipeRepositoryMock.EXPECT().GetRecipesByCategory(suggestedCategory).Return(testdata_fixtures.GetRecipesWithCategory(suggestedCategory))

	// When
	result := s.recipeHandler.GetSuggestion(now)

	// Then
	assert.NotEmpty(t, result.Id)
	assert.NotEmpty(t, result.CreatedAt)
	assert.NotEmpty(t, result.UpdatedAt)
	assert.Equal(t, suggestedCategory, result.Category)
}
