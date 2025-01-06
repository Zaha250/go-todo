package todoModels

type Todo struct {
	//gorm.Model
	Id        TodoId `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoId string

type CreateTodo struct {
	Title string `json:"title"`
}
