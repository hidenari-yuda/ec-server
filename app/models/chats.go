package models

import (
	"log"
	"time"
)

type Chat struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateChat(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at,
		deadline) values(?, ?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetChat(id int) (todo Todo, err error) {
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

func GetChats() (todos []Todo, err error) {
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

func (u *User) GetChatsByUser() (todos []Todo, err error) {

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

func (t *Todo) UpdateChat() error {
	cmd := `update todos set content = ?, user_id =?, deadline = ? where id = ?`

	_, err = Db.Exec(cmd, t.Content, t.UserID, t.Deadline, t.ID)

	return err

}

func (t *Todo) DeleteChat() error {
	cmd := `delete from todos where id =?`

	_, err = Db.Exec(cmd, t.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
