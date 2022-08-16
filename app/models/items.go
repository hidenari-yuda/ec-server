package models

import (
	"log"
	"time"
)

type Item struct {
	ID             int
	UserID         int
	CreatedAt      time.Time
	PhotoURL       string
	Title          string
	Content        string
	CategoryFirst  string
	CategorySecond string
	CategoryThird  string
	Price          int
}

func (u *User) CreateItem(photo_url string, title string, content string, category_first string, category_second string, category_third string, price int) (err error) {
	cmd := `insert into items(
		user_id,
		created_at,
		photo_url,
		title,
		content,
		category_first,
		category_second,
		category_third,
		price) values(?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, time.Now(), photo_url, title, content, category_first, category_second, category_third, price)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetItem(id int) (item Item, err error) {
	cmd := `select id, user_id, created_at, photo_url, title, content, category_first, category_second, category_third, price
	from items where id =?`
	item = Item{}

	err = Db.QueryRow(cmd, id).Scan(
		&item.ID,
		&item.UserID,
		&item.CreatedAt,
		&item.PhotoURL,
		&item.Title,
		&item.Content,
		&item.CategoryFirst,
		&item.CategorySecond,
		&item.CategoryThird,
		&item.Price)

	return item, err
}

func (u *User) GetItemsByOthers() (items []Item, err error) {
	cmd := `select 
	id, 
	user_id,
	created_at,
	photo_url,
	title,
	price from items where user_id != ?`

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
		cmd = `select id, user_id, created_at, photo_url, title, price from items 
		order by created_at desc`

	} else if sortType == "createdat_desc" {
		cmd = `select id, user_id, created_at, photo_url, title, price from items 
		order by created_at asc`
	} else {
		cmd = `select id, user_id, created_at, photo_url, title, price
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
			&item.UserID,
			&item.CreatedAt,
			&item.PhotoURL,
			&item.Title,
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

func GetItemsByFavorites(id int) (items []Item, err error) {
	cmd := `select id, user_id, created_at, photo_url, title, price from items
	where id in (select item_id from items where id =?)`
	rows, err := Db.Query(cmd, id)
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
	cmd := `update items set photo_url = ?, title = ?, content = ?, category_first= ? ,category_second = ?, category_third = ?, price = ? where id = ?`

	_, err = Db.Exec(cmd, i.PhotoURL, i.Title, i.Content, i.CategoryFirst, i.CategorySecond, i.CategoryThird, i.Price, i.ID)

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
