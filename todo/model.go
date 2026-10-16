package todo

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	// Description string `json:"description"`
	IsDone bool `json:"is_done"`
}

type CreateTodoRequest struct {
	Title     string `json:"title" binding:"required"`
	IsDone    bool   `json:"is_done"`
}

type UpdateTodoRequest struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
