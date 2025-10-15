package repository

import (
	"context"
	"danielherschel/home-recipe/pkg/domain"
)

type RecipeBookRepository interface {
	GetRecipeBook(ctx context.Context, user_id string, id string) (*domain.RecipeBook, error)
	SaveRecipeBook(ctx context.Context, user_id string, book *domain.RecipeBook) error
	DeleteRecipeBook(ctx context.Context, user_id string, id string) error
	ListRecipeBooks(ctx context.Context, user_id string) ([]*domain.RecipeBook, error)

	SaveRecipe(ctx context.Context, user_id string, recipe *domain.Recipe) error
	DeleteRecipe(ctx context.Context, user_id string, recipeID string) error
	GetRecipe(ctx context.Context, user_id string, recipeID string) (*domain.Recipe, error)
	ListRecipesInBook(ctx context.Context, user_id string, bookID string) ([]*domain.Recipe, error)
}
