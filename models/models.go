package models

import (
	"todo_server/mydb"
	"fmt"
	"log")

const (
	TODO_WRAPPER_TAB = "todo_wrapper"
	TODO_TAB = "todo"
)

type ToDo struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

type TodoWrapper struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	todos []ToDo
}

type TodoOperation struct {
	IdTodoWrapper int32 `json:"todoWrapperId"`
	Todo ToDo `json:"todo"`
}

func GetAllToDoWrappers() ([]*TodoWrapper, error){
	db := mydb.DB

	query := `SELECT id, title FROM ` + TODO_WRAPPER_TAB
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
		return  nil, err;
	}	
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		panic(err.Error())
		return nil, err
	}

	var todoWrappers []*TodoWrapper

	for rows.Next() {
		var id int32
		var title string
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		todoWrapper := &TodoWrapper {Id: id, Title: title}
		todoWrappers = append(todoWrappers, todoWrapper)

	}

	fmt.Println("ResultGet:", stmt)

    return todoWrappers, nil
}

func GetAllToDosFromTodoWrapper(todoWrapperId int64) ([]*ToDo, error){
	db := mydb.DB

	query := `SELECT id, title, done FROM ` + TODO_TAB + ` WHERE todo_wrapper_id=$1`
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
		return  nil, err;
	}	
	defer stmt.Close()

	rows, err := stmt.Query(todoWrapperId)

	if err != nil {
		panic(err.Error())
		return nil, err
	}

	var todos []*ToDo

	for rows.Next() {
		var id int32
		var title string
		var done bool
		err := rows.Scan(&id, &title, &done)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		todo := &ToDo {Id: id, Title: title, Done: done}
		todos = append(todos, todo)

	}

    return todos, nil
}

func InsertToDoWrapper(todoW TodoWrapper) (*TodoWrapper, error) {
	db := mydb.DB

	query := `INSERT INTO ` + TODO_WRAPPER_TAB + `("title") VALUES ($1) RETURNING id`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
		return  nil, err;
	}	
	defer stmt.Close()

	var todoWrapperId int32
	err = stmt.QueryRow(todoW.Title).Scan(&todoWrapperId)
	if err != nil {
		panic(err.Error())
		return nil, err;
	}

	todoWrapper := &TodoWrapper {Id: todoWrapperId, Title: todoW.Title}

	fmt.Println("resultId:", todoWrapperId)

    return todoWrapper, nil;
}

func InsertToDoInTodoWrappers(todoOp TodoOperation) (*ToDo, error) {
	db := mydb.DB

	query := `INSERT INTO ` + TODO_TAB + `("title", "todo_wrapper_id") VALUES ($1, $2) RETURNING id`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
		return  nil, err;
	}	
	defer stmt.Close()

	var todoId int32
	err = stmt.QueryRow(todoOp.Todo.Title, todoOp.IdTodoWrapper).Scan(&todoId)
	if err != nil {
		panic(err.Error())
		return nil, err;
	}

	todo := &ToDo {Id: todoId, Title: todoOp.Todo.Title, Done: false}

	fmt.Println("resultId:", todoId)

    return todo, nil;
}

func MarkTODOAsDone(todoOperation TodoOperation) error {
	fmt.Printf("CHECK TODO %v AS %v. todoWrapperId: %d / todoId: %d.\n", todoOperation.Todo.Title, todoOperation.Todo.Done, todoOperation.IdTodoWrapper, todoOperation.Todo.Id)
//TODO Alter Done value from a specific ToDo

	db := mydb.DB

	query := `UPDATE ` + TODO_TAB + ` SET done=$1 WHERE id=$2 AND todo_wrapper_id=$3`  

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
		return  err;
	}	
	defer stmt.Close()

	_, err = stmt.Exec(todoOperation.Todo.Done, todoOperation.Todo.Id, todoOperation.IdTodoWrapper)

	if err != nil {
		panic(err.Error())
		return  err;
	}

	return nil
}