package domain

type RecipeBook struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
}

type Recipe struct {
	ID          string   `json:"id"`
	BookID	 	string   `json:"book_id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Instructions string   `json:"instructions"`
}
