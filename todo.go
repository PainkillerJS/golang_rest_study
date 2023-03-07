package todo

type TodoList struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Descripion string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Descripion string `json:"description"`
	Done       bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
