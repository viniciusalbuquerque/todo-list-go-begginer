package models

type ToDo struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

type TodoWrapper struct {
	id int32
	todos []ToDo
}

type TodoOperation struct {
	IdTodoWrapper int32 `json:"todoWrapperId"`
	Todo ToDo `json:"todo"`
}