package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
	Deadline  string
}

func (u *User) CreateTodo(content string, deadline string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at,
		deadline) values(?, ?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now(), deadline)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id , content ,user_id, created_at, deadline
	from todos where id =?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
		&todo.Deadline)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id , content ,user_id, created_at, deadline from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.Deadline)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {

	cmd := `select id, content ,user_id, created_at, deadline from todos where user_id =?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.Deadline)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosBySort(sortType string) (todos []Todo, err error) {
	var (
		cmd string
	)
	if sortType == "createdat_asc" {
		cmd = `select id , content ,user_id, created_at, deadline from todos order by created_at desc`

	} else if sortType == "createdat_desc" {
		cmd = `select id , content ,user_id, created_at, deadline from todos order by created_at asc`

	} else if sortType == "deadline_asc" {
		cmd = `select id , content ,user_id, created_at, deadline from todos order by deadline asc`

	} else {
		cmd = `select id, content ,user_id, created_at, deadline from todos where user_id =?`

	}

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.Deadline)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id =?, deadline = ? where id = ?`

	_, err = Db.Exec(cmd, t.Content, t.UserID, t.Deadline, t.ID)

	return err

}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id =?`

	_, err = Db.Exec(cmd, t.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
