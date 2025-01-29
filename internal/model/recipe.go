package model

import (
	"gororoba/internal/domain"
	"time"
)

type RecipeModel struct {
	Id             string              `json:"id"`
	Title          string              `json:"title"`
	Description    string              `json:"description"`
	Servings       int                 `json:"servings"`
	PrepTime       int                 `json:"prepTime"`
	Slug           string              `json:"slug"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	Ingredients    []domain.Ingredient `json:"ingredients"`
	Category       string              `json:"category"`
	IdAndUpdatedAt string              `json:"id#updatedAt"`
}
