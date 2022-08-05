package models

/*import (
	"log"
	"time"
)

func (u *User) CreateChatGroup(content string) (err error) {
	cmd := `insert into chats (
		content,
		user_id,
		created_at,
		group_id) values(?, ?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now(), cg.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetChatGroup(id int) (chat_group ChatGroup, err error) {
	cmd := `select id, follow_id, follower_id, created_at
	from chat_groups where id =?`
	chat = Chat{}

	err = Db.QueryRow(cmd, id).Scan(
		&chat_groups.ID,
		&chat_groups.Content,
		&chat_groups.UserID,
		&chat_groups.CreatedAt,
		&chat_groups.GroupID)
	return chat_group, err
}

func GetChats() (chats []Chat, err error) {
	cmd := `select id , content ,user_id, created_at, group_id from chats`
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
			&chat.GroupID)
		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}

func (u *User) GetChatsByUser() (chats []Chat, err error) {

	cmd := `select id, content ,user_id, created_at, group_id from chats where user_id =?`

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
			&chat.GroupID)
		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}

func (c *Chat) UpdateChat() error {
	cmd := `update chats set content = ?, user_id =?, where id = ?`

	_, err = Db.Exec(cmd, c.Content, c.UserID, c.ID)

	return err

}

func (c *Chat) DeleteChat() error {
	cmd := `delete from chats where id =?`

	_, err = Db.Exec(cmd, c.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}*/
