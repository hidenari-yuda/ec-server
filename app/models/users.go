package models

import (
	"log"
	"time"
)

type User struct {
	ID         int
	UUID       string
	CreatedAt  time.Time
	Name       string
	NickName   string
	Email      string
	PassWord   string
	IconURL    string
	Phone      string
	Address    string
	Birthday   string
	Item       Item
	Items      []Item
	ChatGroup  ChatGroup
	ChatGroups []ChatGroup
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		created_at,
		name,
		nick_name,
		email,
		password,
		icon_url,
		phone,
		address,
		birthday) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		time.Now(),
		u.Name,
		u.NickName,
		u.Email,
		Encrypt(u.PassWord),
		u.IconURL,
		u.Phone,
		u.Address,
		u.Birthday)

	if err != nil {
		log.Fatalln(err)
	}

	return err

}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, created_at, name, nick_name, email, password, icon_url, phone, address, birthday
	from users where id =?`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.CreatedAt,
		&user.Name,
		&user.NickName,
		&user.Email,
		&user.PassWord,
		&user.IconURL,
		&user.Phone,
		&user.Address,
		&user.Birthday,
	)

	return user, err
}

func GetUsers() (users []User, err error) {
	cmd := `select id, uuid, created_at, name, nick_name, email, password, icon_url, phone, address, birthday
	from users`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.UUID,
			&user.CreatedAt,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.PassWord,
			&user.IconURL,
			&user.Phone,
			&user.Address,
			&user.Birthday,
		)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	rows.Close()

	return users, err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, created_at, name, nick_name, email, password, icon_url, phone, address, birthday
	from users where id =?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.CreatedAt,
		&user.Name,
		&user.NickName,
		&user.Email,
		&user.PassWord,
		&user.IconURL,
		&user.Phone,
		&user.Address,
		&user.Birthday,
	)
	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, created_at, name, nick_name, email, password, icon_url, phone, address, birthday
	from users where email =?`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.CreatedAt,
		&user.Name,
		&user.NickName,
		&user.Email,
		&user.PassWord,
		&user.IconURL,
		&user.Phone,
		&user.Address,
		&user.Birthday,
	)

	return user, err

}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ?, password = ?, icon_url = ?, phone = ?, address = ?, birthday =?
	where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, Encrypt(u.PassWord), u.IconURL, u.Phone, u.Address, u.Birthday, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id =?`

	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions
	(uuid, email, user_id, created_at)  values(?,?,?,?)`

	_, err = Db.Exec(cmd1,
		createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at 
	from sessions where user_id =? and email =?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err

}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at 
	from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid =?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Println(err)
	}
	return err
}
