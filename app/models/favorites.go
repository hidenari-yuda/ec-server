package models

import (
	"database/sql"
	"log"
	"time"
)

type Favorites struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	ItemID    int
	ItemIDs   []int
}

func (u *User) CreateFavorites(item_id int) (err error) {
	cmd := `insert into favoritess(
		user_id,
		created_at,
		item_id) values(?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, time.Now(), item_id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetFavorites(id int) (favorite Favorites, err error) {
	cmd := `SELECT id, user_id, created_at, item_id
	FROM favorites WHERE id =?`
	favorite = Favorites{}

	err = Db.QueryRow(cmd, id).Scan(
		&favorite.ID,
		&favorite.UserID,
		&favorite.CreatedAt,
		&favorite.ItemID)

	return favorite, err
}

func GetItemsByFavorites(ids []int) (items []Item, err error) {
	var (
		cmd  string
		rows *sql.Rows
		id   int
	)

	for i := 0; i > len(ids); i++ {
		id = ids[i]
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
	WHERE id = ?`
		rows, err = Db.Query(cmd, id)
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

	}

	return items, err
}

func (f *Favorites) DeleteFavorites() error {
	cmd := `delete FROM favorites WHERE id =?`

	_, err = Db.Exec(cmd, f.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
