package todoModels

type Todo struct {
	//gorm.Model
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type CreateTodo struct {
	Title string `json:"title"`
}
