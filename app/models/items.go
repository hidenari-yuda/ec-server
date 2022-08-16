package models

import (
	"log"
	"time"
)

type Item struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	PhotoURL  string
	Title     string
	Content   string
	Category  string
	Price     int
}

func (u *User) CreateItem(photo_url string, title string, content string, category string, price int) (err error) {
	cmd := `insert into items(
		user_id,
		created_at,
		photo_url,
		title,
		content,
		category,
		price) values(?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, time.Now(), photo_url, title, content, category, price)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetItem(id int) (item Item, err error) {
	cmd := `select id, user_id, created_at, photo_url, title, content, category, price
	from items where id =?`
	item = Item{}

	err = Db.QueryRow(cmd, id).Scan(
		&item.ID,
		&item.UserID,
		&item.CreatedAt,
		&item.PhotoURL,
		&item.Title,
		&item.Content,
		&item.Category,
		&item.Price)

	return item, err
}

func GetAllItems() (items []Item, err error) {
	cmd := `select 
	id, 
	user_id,
	created_at,
	photo_url,
	title,
	content,
	category,
	price from items`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var item Item
		err = rows.Scan(
			&item.ID,
			&item.UserID,
			&item.CreatedAt,
			&item.PhotoURL,
			&item.Title,
			&item.Content,
			&item.Category,
			&item.Price)
		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, item)
	}
	rows.Close()

	return items, err
}

func (u *User) GetItemsByUser() (items []Item, err error) {

	cmd := `select
	id, 
	user_id,
	created_at,
	photo_url,
	title,
	content,
	category,
	price  from items where user_id =?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var item Item
		err = rows.Scan(
			&item.ID,
			&item.UserID,
			&item.CreatedAt,
			&item.PhotoURL,
			&item.Title,
			&item.Content,
			&item.Category,
			&item.Price,
		)

		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, item)
	}
	rows.Close()

	return items, err
}

func (u *User) GetItemsBySort(sortType string) (items []Item, err error) {
	var (
		cmd string
	)
	if sortType == "createdat_asc" {
		cmd = `select id, user_id, created_at, photo_url, title, content, category, price from items 
		order by created_at desc`

	} else if sortType == "createdat_desc" {
		cmd = `select id, user_id, created_at, photo_url, title, content, category, price from items 
		order by created_at asc`
	} else {
		cmd = `select id, user_id, created_at, photo_url, title, content, category, price
		from items where user_id =?`
	}

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var item Item
		err = rows.Scan(
			&item.ID,
			&item.ID,
			&item.UserID,
			&item.CreatedAt,
			&item.PhotoURL,
			&item.Title,
			&item.Content,
			&item.Category,
			&item.Price,
		)

		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, item)
	}
	rows.Close()

	return items, err
}

func (i *Item) UpdateItem() error {
	cmd := `update items set photo_url = ?, title = ?, content = ?, category = ?, price = ? where id = ?`

	_, err = Db.Exec(cmd, i.PhotoURL, i.Title, i.Content, i.Category, i.Price, i.ID)

	return err

}

func (i *Item) DeleteItem() error {
	cmd := `delete from items where id =?`

	_, err = Db.Exec(cmd, i.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
