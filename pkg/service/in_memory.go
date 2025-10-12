package service

import (
	"context"
	"errors"
	"sync"

	"danielherschel/home-recipe/pkg/domain"
)

func NewInMemoryService() *InMemoryService {
	return &InMemoryService{
		books:   make(map[string]*domain.RecipeBook),
		recipes: make(map[string]*domain.Recipe),
	}
}

type InMemoryService struct {
	mu      sync.RWMutex
	books   map[string]*domain.RecipeBook
	recipes map[string]*domain.Recipe
}

var (
	ErrNotFound = errors.New("not found")
)

// Ensure InMemoryService implements RecipeBookService
var _ RecipeBookService = (*InMemoryService)(nil)

func (s *InMemoryService) GetRecipeBook(ctx context.Context, user_id string, id string) (*domain.RecipeBook, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[id]
	if !ok || b.UserID != user_id {
		return nil, ErrNotFound
	}
	// Return a copy to avoid accidental external mutation
	copy := *b
	return &copy, nil
}

func (s *InMemoryService) SaveRecipeBook(ctx context.Context, user_id string, book *domain.RecipeBook) error {
	if book == nil || book.ID == "" {
		return errors.New("invalid book")
	}
	// enforce user_id
	if book.UserID == "" {
		book.UserID = user_id
	} else if book.UserID != user_id {
		return errors.New("user mismatch")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	// store a copy
	c := *book
	s.books[book.ID] = &c
	return nil
}

func (s *InMemoryService) DeleteRecipeBook(ctx context.Context, user_id string, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.books[id]
	if !ok || b.UserID != user_id {
		return ErrNotFound
	}
	delete(s.books, id)
	// delete associated recipes
	for rid, r := range s.recipes {
		if r.BookID == id {
			delete(s.recipes, rid)
		}
	}
	return nil
}

func (s *InMemoryService) ListRecipeBooks(ctx context.Context, user_id string) ([]*domain.RecipeBook, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*domain.RecipeBook, 0, len(s.books))
	for _, b := range s.books {
		if b.UserID != user_id {
			continue
		}
		c := *b
		out = append(out, &c)
	}
	return out, nil
}

func (s *InMemoryService) SaveRecipe(ctx context.Context, user_id string, recipe *domain.Recipe) error {
	if recipe == nil || recipe.ID == "" || recipe.BookID == "" {
		return errors.New("invalid recipe")
	}
	// enforce user_id
	if recipe.UserID == "" {
		recipe.UserID = user_id
	} else if recipe.UserID != user_id {
		return errors.New("user mismatch")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	// ensure book exists and belongs to user
	b, ok := s.books[recipe.BookID]
	if !ok || b.UserID != user_id {
		return errors.New("book not found")
	}
	c := *recipe
	// copy ingredients slice
	if recipe.Ingredients != nil {
		c.Ingredients = make([]string, len(recipe.Ingredients))
		copy(c.Ingredients, recipe.Ingredients)
	}
	s.recipes[recipe.ID] = &c
	return nil
}

func (s *InMemoryService) DeleteRecipe(ctx context.Context, user_id string, recipeID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.recipes[recipeID]
	if !ok || r.UserID != user_id {
		return ErrNotFound
	}
	delete(s.recipes, recipeID)
	return nil
}

func (s *InMemoryService) GetRecipe(ctx context.Context, user_id string, recipeID string) (*domain.Recipe, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.recipes[recipeID]
	if !ok || r.UserID != user_id {
		return nil, ErrNotFound
	}
	c := *r
	if r.Ingredients != nil {
		c.Ingredients = make([]string, len(r.Ingredients))
		copy(c.Ingredients, r.Ingredients)
	}
	return &c, nil
}

func (s *InMemoryService) ListRecipesInBook(ctx context.Context, user_id string, bookID string) ([]*domain.Recipe, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[bookID]
	if !ok || b.UserID != user_id {
		return nil, ErrNotFound
	}
	out := make([]*domain.Recipe, 0)
	for _, r := range s.recipes {
		if r.BookID == bookID && r.UserID == user_id {
			c := *r
			if r.Ingredients != nil {
				c.Ingredients = make([]string, len(r.Ingredients))
				copy(c.Ingredients, r.Ingredients)
			}
			out = append(out, &c)
		}
	}
	return out, nil
}
