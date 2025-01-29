package testdata_fixtures

import (
	"time"

	"gororoba/internal/domain"
)

func GetRecipesWithCategory(category string) []domain.Recipe {
	return []domain.Recipe{
		{
			Id:          "1",
			Title:       "Test Recipe",
			Description: "This is a test recipe",
			Servings:    4,
			PrepTime:    30,
			Slug:        "test-recipe",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Ingredients: []domain.Ingredient{
				{
					Name:        "Test Ingredient",
					Quantity:    "1",
					MeasureUnit: domain.MilliLiter,
				},
			},
			Category: category,
		},
		{
			Id:          "2",
			Title:       "Test Recipe 2",
			Description: "This is a test recipe",
			Servings:    6,
			PrepTime:    120,
			Slug:        "test-recipe-2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Ingredients: []domain.Ingredient{
				{
					Name:        "Test Ingredient",
					Quantity:    "1/2",
					MeasureUnit: domain.MilliLiter,
				},
			},
			Category: category,
		},
	}
}
