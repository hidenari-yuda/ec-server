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
	GroupID   int
	UserName  string
	PhotoURL  string
}

func (u *User) CreateChat(content string, group_id int) (err error) {
	cmd := `insert into messages (
		content,
		user_id,
		created_at,
		group_id,
		user_name) values(?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now(), group_id, u.Name)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetChat(id int) (chat Chat, err error) {
	cmd := `select id , content ,user_id, created_at, group_id, user_name
	from messages where id =?`
	chat = Chat{}

	err = Db.QueryRow(cmd, id).Scan(
		&chat.ID,
		&chat.Content,
		&chat.UserID,
		&chat.CreatedAt,
		&chat.GroupID,
		&chat.UserName)
	return chat, err
}

func GetChats() (chats []Chat, err error) {
	cmd := `select id , content ,user_id, created_at, group_id, user_name from messages`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var chat Chat
		err = rows.Scan(
			&chat.ID,
			&chat.Content,
			&chat.UserID,
			&chat.CreatedAt,
			&chat.GroupID,
			&chat.UserName)
		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}

func (u *User) GetChatsByUser() (chats []Chat, err error) {

	cmd := `select id, content ,user_id, created_at, group_id, user_name from messages where user_id =?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var chat Chat
		err = rows.Scan(
			&chat.ID,
			&chat.Content,
			&chat.UserID,
			&chat.CreatedAt,
			&chat.GroupID,
			&chat.UserName)
		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}

func (u *User) GetChatsByGroup(group_id int) (chats []Chat, err error) {

	cmd := `select id, content ,user_id, created_at, group_id, user_name from messages where group_id =?`

	rows, err := Db.Query(cmd, group_id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var chat Chat
		err = rows.Scan(
			&chat.ID,
			&chat.Content,
			&chat.UserID,
			&chat.CreatedAt,
			&chat.GroupID,
			&chat.UserName)
		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}

func (c *Chat) UpdateChat() error {
	cmd := `update messages set content = ?, user_id =?, where id = ?`

	_, err = Db.Exec(cmd, c.Content, c.UserID, c.ID)

	return err

}

func (c *Chat) DeleteChat() error {
	cmd := `delete from messages where id =?`

	_, err = Db.Exec(cmd, c.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
