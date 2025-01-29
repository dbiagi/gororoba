package converter_test

import (
	"testing"
	"time"

	"gororoba/internal/converter"
	"gororoba/internal/domain"
	"gororoba/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestToRecipeModel(t *testing.T) {
	// Given
	createdAt := time.Now()
	updatedAt := createdAt.Add(24 * time.Hour)
	domainRecipe := domain.Recipe{
		Id:          "1",
		Title:       "Test Recipe",
		Description: "Test Description",
		Servings:    4,
		PrepTime:    30,
		Slug:        "test-recipe",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Ingredients: []domain.Ingredient{},
		Category:    "Test Category",
	}

	// When
	modelRecipe := converter.ToRecipeModel(domainRecipe)

	// Then
	assert.Equal(t, domainRecipe.Id, modelRecipe.Id)
	assert.Equal(t, domainRecipe.Title, modelRecipe.Title)
	assert.Equal(t, domainRecipe.Description, modelRecipe.Description)
	assert.Equal(t, domainRecipe.Servings, modelRecipe.Servings)
	assert.Equal(t, domainRecipe.PrepTime, modelRecipe.PrepTime)
	assert.Equal(t, domainRecipe.Slug, modelRecipe.Slug)
	assert.Equal(t, domainRecipe.CreatedAt, modelRecipe.CreatedAt)
	assert.Equal(t, domainRecipe.UpdatedAt, modelRecipe.UpdatedAt)
	assert.Equal(t, domainRecipe.Ingredients, modelRecipe.Ingredients)
	assert.Equal(t, domainRecipe.Category, modelRecipe.Category)
	assert.Equal(t, domainRecipe.Id+"#"+domainRecipe.UpdatedAt.Format(time.RFC3339), modelRecipe.IdAndUpdatedAt)
}

func TestToRecipeDomain(t *testing.T) {
	// Given
	createdAt := time.Now()
	updatedAt := createdAt.Add(24 * time.Hour)
	modelRecipe := model.RecipeModel{
		Id:          "1",
		Title:       "Test Recipe",
		Description: "Test Description",
		Servings:    4,
		PrepTime:    30,
		Slug:        "test-recipe",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Ingredients: []domain.Ingredient{},
		Category:    "Test Category",
	}

	// When
	domainRecipe := converter.ToRecipeDomain(modelRecipe)

	// Then
	assert.Equal(t, modelRecipe.Id, domainRecipe.Id)
	assert.Equal(t, modelRecipe.Title, domainRecipe.Title)
	assert.Equal(t, modelRecipe.Description, domainRecipe.Description)
	assert.Equal(t, modelRecipe.Servings, domainRecipe.Servings)
	assert.Equal(t, modelRecipe.PrepTime, domainRecipe.PrepTime)
	assert.Equal(t, modelRecipe.Slug, domainRecipe.Slug)
	assert.Equal(t, modelRecipe.CreatedAt, domainRecipe.CreatedAt)
	assert.Equal(t, modelRecipe.UpdatedAt, domainRecipe.UpdatedAt)
	assert.Equal(t, modelRecipe.Ingredients, domainRecipe.Ingredients)
	assert.Equal(t, modelRecipe.Category, domainRecipe.Category)
}
