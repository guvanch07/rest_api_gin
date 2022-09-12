package data

type Todo struct {
	ID        string `json:"id"`
	Item      string `json: "title"`
	Completed bool   `json: "completed"`
}

var Todos = []Todo{
	{ID: "1", Item: "clean room", Completed: false},
	{ID: "2", Item: "read book", Completed: true},
	{ID: "3", Item: "cook dinner", Completed: false},
}
