package handler

import (
	"math/rand/v2"
	"time"

	"gororoba/internal/converter"
	"gororoba/internal/domain"
	"gororoba/internal/repository"

	"github.com/google/uuid"
)

type RecipesHandlerInterface interface {
	GetRecipesByCategory(category string) []domain.Recipe
	CreateRecipe(r *domain.Recipe) *domain.Recipe
	GetSuggestion(suggestionRequestedAt time.Time) domain.Recipe
}

type RecipesHandler struct {
	RecipeRepository  repository.RecipeRepositoryInterface
	SuggestionHandler SuggestionHandlerInterface
}

func NewRecipesHandler(r repository.RecipeRepositoryInterface, sh SuggestionHandlerInterface) RecipesHandler {
	return RecipesHandler{
		RecipeRepository:  r,
		SuggestionHandler: sh,
	}
}

func (h RecipesHandler) GetRecipesByCategory(category string) []domain.Recipe {
	return h.RecipeRepository.GetRecipesByCategory(category)
}

func (h RecipesHandler) GetSuggestion(suggestionRequestedAt time.Time) domain.Recipe {
	suggestedCategory := h.SuggestionHandler.GetSuggestedCategoryByTime(suggestionRequestedAt)

	recipes := h.RecipeRepository.GetRecipesByCategory(suggestedCategory)

	len := len(recipes)
	i := rand.IntN(len - 1)

	return recipes[i]
}

func (h RecipesHandler) CreateRecipe(r *domain.Recipe) *domain.Recipe {
	r.Id = uuid.New().String()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	m := converter.ToRecipeModel(*r)
	h.RecipeRepository.CreateRecipe(m)

	return r
}
