package controller

import (
	"gororoba/internal/handler"
	"net/http"
	"time"
)

type RecipesController struct {
	RecipesHandler handler.RecipesHandlerInterface
}

func NewRecipesController(h handler.RecipesHandlerInterface) RecipesController {
	return RecipesController{RecipesHandler: h}
}

func (rc *RecipesController) GetRecipesByCategory(w http.ResponseWriter, r *http.Request) HttpResponse {
	category := r.URL.Query().Get("category")
	recipes := rc.RecipesHandler.GetRecipesByCategory(category)
	return HttpResponse{Body: recipes}
}

func (rc *RecipesController) GetSuggestion(w http.ResponseWriter, r *http.Request) HttpResponse {
	recipes := rc.RecipesHandler.GetSuggestion(time.Now())
	return HttpResponse{Body: recipes}
}
