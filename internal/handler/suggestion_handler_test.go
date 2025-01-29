package handler_test

import (
	"fmt"
	"gororoba/internal/domain"
	"gororoba/internal/handler"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type _suggestionHandlerSetup struct {
	suggestionHandler *handler.SuggestionHandler
}

func suggestionHandlerSetup() _suggestionHandlerSetup {
	sh := handler.NewSuggestionHandler()

	return _suggestionHandlerSetup{
		suggestionHandler: &sh,
	}
}

func TestGetSuggestion(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")

	if err != nil {
		t.Error(err)
	}

	// Given
	s := suggestionHandlerSetup()
	cases := []struct {
		time             time.Time
		expectedCategory string
	}{
		{date.Add(time.Hour * 6), string(domain.RecipeCategoryBreakfast)},
		{date.Add(time.Hour * 12), string(domain.RecipeCategoryMainCourse)},
		{date.Add(time.Hour * 18), string(domain.RecipeCategoryDessert)},
		{date, string(domain.RecipeCategorySnack)},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("given the hour %d should suggest category %s", c.time.Hour(), c.expectedCategory), func(t *testing.T) {
			// When
			result := s.suggestionHandler.GetSuggestedCategoryByTime(c.time)

			// Then
			assert.Equal(t, c.expectedCategory, result)
		})
	}
}
