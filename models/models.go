package models

import ("todo_server/mydb")

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

func GetAllToDoWrappers() ([]*TodoWrapper, error){
	rows, err := mydb.DB.Query("SELECT * FROM todo_wrapper")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return nil, nil
}