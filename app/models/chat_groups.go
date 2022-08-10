package models

import (
	"log"
	"time"
)

type ChatGroup struct {
	ID         int
	UserID     int
	CreatedAt  time.Time
	ChatMember string
	ChatName   string
	Chat       []Chat
}

func (u *User) CreateChatGroup(ChatMember interface{}, ChatName string) (err error) {
	cmd := `insert into chatgroups (
		user_id,
		created_at,
		chat_member,
		chat_name) values(?, ?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, time.Now(), ChatMember, ChatName)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetChatGroup(id int) (chat_group ChatGroup, err error) {
	cmd := `select id, user_id, created_at, chat_member, chat_name
	from chatgroups where id =?`
	chat_group = ChatGroup{}

	err = Db.QueryRow(cmd, id).Scan(
		&chat_group.ID,
		&chat_group.UserID,
		&chat_group.CreatedAt,
		&chat_group.ChatMember,
		&chat_group.ChatName)
	return chat_group, err
}

func (u *User) GetChatGroupByID(id int) (chat_group ChatGroup, err error) {
	cmd := `select id, user_id, created_at, chat_member, chat_name
	from chatgroups where id =?`
	chat_group = ChatGroup{}

	err = Db.QueryRow(cmd, id).Scan(
		&chat_group.ID,
		&chat_group.UserID,
		&chat_group.CreatedAt,
		&chat_group.ChatMember,
		&chat_group.ChatName)
	return chat_group, err
}

func GetChatGroups() (chatgroups []ChatGroup, err error) {
	cmd := `select id ,user_id, created_at, chat_member, chat_name from chatgroups`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var chatgroup ChatGroup
		err = rows.Scan(
			&chatgroup.ID,
			&chatgroup.UserID,
			&chatgroup.CreatedAt,
			&chatgroup.ChatMember,
			&chatgroup.ChatName)

		if err != nil {
			log.Fatalln(err)
		}
		chatgroups = append(chatgroups, chatgroup)
	}
	rows.Close()

	return chatgroups, err
}

func (u *User) GetChatGroupsByUser() (chatgroups []ChatGroup, err error) {

	cmd := `select id, user_id, created_at, chat_member, chat_name from chatgroups where user_id =? or chat_member =?`

	rows, err := Db.Query(cmd, u.ID, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var chatgroup ChatGroup
		err = rows.Scan(
			&chatgroup.ID,
			&chatgroup.UserID,
			&chatgroup.CreatedAt,
			&chatgroup.ChatMember,
			&chatgroup.ChatName)
		if err != nil {
			log.Fatalln(err)
		}
		chatgroups = append(chatgroups, chatgroup)
	}
	rows.Close()

	return chatgroups, err
}

/*func (cg *ChatGroup) UpdateChatGroup() error {
	cmd := `update chatgroups set chat_name = ?, where id = ?`

	_, err = Db.Exec(cmd, cg.ChatName, cg.ID)

	return err

}*/

func (cg *ChatGroup) DeleteChatGroup() error {
	cmd := `delete from chatgroups where id =?`

	_, err = Db.Exec(cmd, cg.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
