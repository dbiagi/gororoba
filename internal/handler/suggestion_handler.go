package handler

import (
	"gororoba/internal/domain"
	"time"
)

type SuggestionHandlerInterface interface {
	GetSuggestedCategoryByTime(t time.Time) string
}

type SuggestionHandler struct {
}

func NewSuggestionHandler() SuggestionHandler {
	return SuggestionHandler{}
}

func (sh SuggestionHandler) GetSuggestedCategoryByTime(t time.Time) string {
	switch {
	case isMorning(t):
		return string(domain.RecipeCategoryBreakfast)
	case isAfternoon(t):
		return string(domain.RecipeCategoryMainCourse)
	case isNight(t):
		return string(domain.RecipeCategoryDessert)
	default:
		return string(domain.RecipeCategorySnack)
	}
}

func isMorning(t time.Time) bool {
	return t.Hour() >= 6 && t.Hour() < 12
}

func isAfternoon(t time.Time) bool {
	return t.Hour() >= 12 && t.Hour() < 18
}

func isNight(t time.Time) bool {
	return t.Hour() >= 18 && t.Hour() < 24
}

//go:generate mockgen -destination=../testdata/mocks/suggestion_handler_mock.go -package=mocks gororoba/internal/handler SuggestionHandlerInterface
