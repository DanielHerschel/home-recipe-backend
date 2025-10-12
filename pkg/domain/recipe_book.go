package domain

type RecipeBook struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	UserID  string   `json:"user_id"`
}

type Recipe struct {
	ID          string   `json:"id"`
	UserID	  string   `json:"user_id"`
	BookID	 	string   `json:"book_id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Instructions string   `json:"instructions"`
}
