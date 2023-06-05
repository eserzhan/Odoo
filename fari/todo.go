package todo

type Todo_lists struct {
	Id int `json:"id" bd: "id"`
	Title string `json:"title" bd: "title" binding: "required"`
	Description string `json:"description" bd: "description"`
}

type Users_lists struct {
	Id int 
	UserId int 
	ListId int 
}

type UpdateTodoLists struct {
	Title *string `json:"title" bd: "title" `
	Description *string `json:"description" bd: "description"`
}

type Todo_items struct {
	Id int `json:"id" bd: "id"`
	Title string `json:"title" bd: "title" binding: "required"`
	Description string `json:"description" bd: "description"`
	Done bool `json:"done" bd: "done"`
}

type Lists_items struct {
	Id int 
	ItemId int 
	ListId int 
}

type UpdateTodoItems struct {
	Title *string `json:"title" bd: "title"`
	Description *string `json:"description" bd: "description"`
	Done *bool `json:"done" bd: "done"`
}