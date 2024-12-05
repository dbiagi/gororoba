package converter

import (
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/dbiagi/gororoba/src/model"
	"time"
)

func ToRecipeModel(r domain.Recipe) model.RecipeModel {
	return model.RecipeModel{
		Id:             r.Id,
		Title:          r.Title,
		Description:    r.Description,
		Servings:       r.Servings,
		PrepTime:       r.PrepTime,
		Slug:           r.Slug,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
		Ingredients:    r.Ingredients,
		Category:       r.Category,
		IdAndUpdatedAt: r.Id + "#" + r.UpdatedAt.Format(time.RFC3339),
	}
}

func ToRecipeDomain(r model.RecipeModel) domain.Recipe {
	return domain.Recipe{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		Servings:    r.Servings,
		PrepTime:    r.PrepTime,
		Slug:        r.Slug,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		Ingredients: r.Ingredients,
		Category:    r.Category,
	}
}