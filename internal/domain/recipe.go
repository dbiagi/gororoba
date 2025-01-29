package domain

import (
	"time"
)

type RecipeCategory string

const (
	RecipeCategoryBreakfast  RecipeCategory = "breakfast"
	RecipeCategoryMainCourse RecipeCategory = "main_course"
	RecipeCategorySnack      RecipeCategory = "snack"
	RecipeCategoryDessert    RecipeCategory = "dessert"
)

// PK = Category, SK = Id#UpdatedAt
// GSI = Id
// GSI = Slug
type Recipe struct {
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Servings    int          `json:"servings"`
	PrepTime    int          `json:"prepTime"`
	Slug        string       `json:"slug"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Ingredients []Ingredient `json:"ingredients"`
	Category    string       `json:"category"`
}

type Ingredient struct {
	Name        string
	Quantity    string
	MeasureUnit MeasureUnit
}
