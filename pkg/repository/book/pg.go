package book

import (
	"context"
	"danielherschel/home-recipe/pkg/domain"
	"errors"

	db "danielherschel/home-recipe/pkg/database"
)

func NewRecipeBookRepository() *PostgresRecipeBookRepository {
	return &PostgresRecipeBookRepository{}
}

type PostgresRecipeBookRepository struct {
	DB *db.PGDatabase
}

var _ RecipeBookRepository = (*PostgresRecipeBookRepository)(nil)

func (r *PostgresRecipeBookRepository) GetRecipeBook(ctx context.Context, user_id string, id string) (*domain.RecipeBook, error) {
	row := r.DB.Conn.QueryRow(ctx, "SELECT id, title, user_id FROM recipe_books WHERE id=$1 AND user_id=$2", id, user_id)

	recipeBook := &domain.RecipeBook{}
	err := row.Scan(&recipeBook.ID, &recipeBook.Title, &recipeBook.UserID)
	if err != nil {
		return nil, err
	}
	return recipeBook, nil
}

func (r *PostgresRecipeBookRepository) SaveRecipeBook(ctx context.Context, user_id string, book *domain.RecipeBook) error {
	if book.UserID == "" {
		book.UserID = user_id
	} else if book.UserID != user_id {
		return errors.New("user mismatch")
	}

	_, err := r.DB.Conn.Exec(ctx, `
		INSERT INTO recipe_books (id, title, user_id) 
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET 
			title = EXCLUDED.title
		`, book.ID, book.Title, book.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRecipeBookRepository) DeleteRecipeBook(ctx context.Context, user_id string, id string) error {
	_, err := r.DB.Conn.Exec(ctx, "DELETE FROM recipe_books WHERE id=$1 AND user_id=$2", id, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRecipeBookRepository) ListRecipeBooks(ctx context.Context, user_id string) ([]*domain.RecipeBook, error) {
	rows, err := r.DB.Conn.Query(ctx, "SELECT id, title, user_id FROM recipe_books WHERE user_id=$1", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*domain.RecipeBook
	for rows.Next() {
		book := &domain.RecipeBook{}
		if err := rows.Scan(&book.ID, &book.Title, &book.UserID); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *PostgresRecipeBookRepository) SaveRecipe(ctx context.Context, user_id string, recipe *domain.Recipe) error {
	if recipe.UserID == "" {
		recipe.UserID = user_id
	} else if recipe.UserID != user_id {
		return errors.New("user mismatch")
	}

	// ensure the book belongs to the user
	var bookUser string
	err := r.DB.Conn.QueryRow(ctx, "SELECT user_id FROM recipe_books WHERE id=$1", recipe.BookID).Scan(&bookUser)
	if err != nil {
		return err
	}
	if bookUser != user_id {
		return errors.New("book not found or not owned by user")
	}

	_, err = r.DB.Conn.Exec(ctx, `
		INSERT INTO recipes (id, user_id, book_id, title, ingredients, instructions) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET 
			title = EXCLUDED.title,
			ingredients = EXCLUDED.ingredients,
			instructions = EXCLUDED.instructions
		`, recipe.ID, recipe.UserID, recipe.BookID, recipe.Title, recipe.Ingredients, recipe.Instructions)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRecipeBookRepository) DeleteRecipe(ctx context.Context, user_id string, recipeID string) error {
	_, err := r.DB.Conn.Exec(ctx, "DELETE FROM recipes WHERE id=$1 AND user_id=$2", recipeID, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRecipeBookRepository) GetRecipe(ctx context.Context, user_id string, recipeID string) (*domain.Recipe, error) {
	row := r.DB.Conn.QueryRow(ctx, "SELECT id, user_id, book_id, title, ingredients, instructions FROM recipes WHERE id=$1 AND user_id=$2", recipeID, user_id)

	recipe := &domain.Recipe{}
	err := row.Scan(&recipe.ID, &recipe.UserID, &recipe.BookID, &recipe.Title, &recipe.Ingredients, &recipe.Instructions)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *PostgresRecipeBookRepository) ListRecipesInBook(ctx context.Context, user_id string, bookID string) ([]*domain.Recipe, error) {
	rows, err := r.DB.Conn.Query(ctx, "SELECT id, user_id, book_id, title, ingredients, instructions FROM recipes WHERE book_id=$1 AND user_id=$2", bookID, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*domain.Recipe
	for rows.Next() {
		recipe := &domain.Recipe{}
		if err := rows.Scan(&recipe.ID, &recipe.UserID, &recipe.BookID, &recipe.Title, &recipe.Ingredients, &recipe.Instructions); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}
