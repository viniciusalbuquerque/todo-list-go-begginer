package models

import ("todo_server/mydb")

type ToDo struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

type TodoWrapper struct {
	id int32 `json:"id"`
	Title string `json:"title"`
	todos []ToDo
}

type TodoOperation struct {
	IdTodoWrapper int32 `json:"todoWrapperId"`
	Todo ToDo `json:"todo"`
}

func GetAllToDoWrappers() ([]*TodoWrapper, error){
	db := mydb.DB
	rows, err := db.Query("SELECT * FROM todowrapper")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return nil, nil
}

func InsertToDoWrapper(todoW TodoWrapper) error {
	db := mydb.DB
	stmtIns, err := db.Prepare("INSERT INTO todowrapper (title) VALUES  ? ") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return err;
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(todoW.Title)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

    return err;
}

func InsertToDoInTodoWrappers(todoOp TodoOperation) error {
	db := mydb.DB
	stmtIns, err := db.Prepare("INSERT INTO todo VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return err;
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec("title", todoOp.Todo.Title)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

    return err;
}