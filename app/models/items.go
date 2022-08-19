package models

import (
	"database/sql"
	"fmt"
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
	cmd := `SELECT id, user_id, created_at, photo_url, title, content, category_first, category_second, category_third, price
	FROM items WHERE id =?`
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

func GetAllItems() (items []Item, err error) {
	cmd := `SELECT 
	id, 
	user_id,
	created_at,
	photo_url,
	title,
	price FROM items`

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
			&item.Price)
		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, item)
	}
	rows.Close()

	return items, err
}

func (u *User) GetItemsByOthers() (items []Item, err error) {
	cmd := `SELECT 
	id, 
	user_id,
	created_at,
	photo_url,
	title,
	price FROM items WHERE user_id != ?`

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

	cmd := `SELECT
	id, 
	user_id,
	created_at,
	photo_url,
	title,
	price  FROM items WHERE user_id =?`

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
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
		order by created_at desc`

	} else if sortType == "createdat_desc" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
			order by created_at asc`
	} else {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price
				FROM items WHERE user_id =?`
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

func (u *User) GetItemsByCategory(category_first string, category_second string, category_third string) (items []Item, err error) {
	var (
		cmd  string
		rows *sql.Rows
	)

	if category_first != "99" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
				WHERE user_id = ? AND category_first = ? OR category_second = ? OR category_third = ?`
		rows, err = Db.Query(cmd, u.ID, category_first, category_second, category_third)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
				WHERE user_id = ?`
		rows, err = Db.Query(cmd, u.ID)
		if err != nil {
			log.Fatalln(err)
		}
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

func (u *User) GetItemsByFreeWords(freewords string, category_first string, category_second string, category_third string, price_min int, price_max int) (items []Item, err error) {
	var (
		cmd  string
		rows *sql.Rows
	)
	//フリーワード有り、カテゴリ有り、価格有り
	if freewords != "" && category_first != "99" || category_second != "99" || category_third != "99" && price_min != 0 || price_max != 100000000 {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
				WHERE user_id = ? AND title like ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?) AND price >= ? AND price <= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			price_min, price_max,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワード有り、カテゴリ有り、価格なし
	} else if freewords != "" && category_first != "99" || category_second != "99" || category_third != "99" && price_min == 0 && price_max == 100000000 {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ? AND title like ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?)`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワード有り、カテゴリなし、価格有り
	} else if freewords != "" && (category_first == "99" && category_second == "99" && category_third == "99") && (price_min != 0 || price_max != 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ? AND title like ? AND price >= ? AND price <= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
			price_min, price_max,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワード有り、カテゴリなし、価格なし
	} else if freewords != "" && (category_first == "99" && category_second == "99" && category_third == "99") && (price_min == 0 && price_max == 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ? AND title like ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
		)
		if err != nil {
			log.Fatalln(err)
		}

		//フリーワードなし、カテゴリ有り、価格有り
	} else if freewords == "" && (category_first != "99" || category_second != "99" || category_third != "99") && (price_min != 0 || price_max != 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?) AND price >= ? AND price <= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			price_min, price_max,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワードなし、カテゴリ有り、価格なし
	} else if freewords == "" && (category_first != "99" || category_second != "99" || category_third != "99") && (price_min == 0 && price_max == 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?)`

		rows, err = Db.Query(
			cmd,
			u.ID,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
		)
		if err != nil {
			log.Fatalln(err)
		}

		//フリーワードなし、カテゴリなし、価格あり
	} else if freewords == "" && (category_first == "99" && category_second == "99" && category_third == "99") && (price_min != 0 || price_max != 100000000) {
		fmt.Println("price")
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ? AND price <= ? AND price >= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			price_min, price_max,
		)
	} else {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id = ?`
		rows, err = Db.Query(
			cmd,
			u.ID,
		)
		if err != nil {
			log.Fatalln(err)
		}
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

func (u *User) GetOnsaleBySort(sortType string) (items []Item, err error) {
	var cmd string
	if sortType == "createdat_asc" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
		order by created_at desc`
	} else if sortType == "createdat_desc" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
		order by created_at asc`
	} else if sortType == "price_high" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		order by price desc`
	} else if sortType == "price_low" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		order by price asc`
	} else {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price
		FROM items WHERE user_id != ?`
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

func (u *User) GetOnsaleByCategory(category_first string, category_second string, category_third string) (items []Item, err error) {
	var (
		cmd  string
		rows *sql.Rows
	)
	if category_first != "99" {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items 
		WHERE user_id != ? AND category_first = ? OR category_second = ? OR category_third = ?`

		rows, err = Db.Query(cmd, u.ID, category_first, category_second, category_third)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ?`

		rows, err = Db.Query(cmd, u.ID)
		if err != nil {
			log.Fatalln(err)
		}
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

func (u *User) GetOnsaleByFreeWords(freewords string, category_first string, category_second string, category_third string, price_min int, price_max int) (items []Item, err error) {
	var (
		cmd  string
		rows *sql.Rows
	)
	//フリーワード有り、カテゴリ有り、価格有り
	if freewords != "" && category_first != "99" || category_second != "99" || category_third != "99" && price_min != 0 || price_max != 100000000 {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND title like ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?) AND price >= ? AND price <= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			price_min, price_max,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワード有り、カテゴリ有り、価格なし
	} else if freewords != "" && category_first != "99" || category_second != "99" || category_third != "99" && price_min == 0 && price_max == 100000000 {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND title like ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?)`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワード有り、カテゴリなし、価格有り
	} else if freewords != "" && (category_first == "99" && category_second == "99" && category_third == "99") && (price_min != 0 || price_max != 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND title like ? AND price >= ? AND price <= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
			price_min, price_max,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワード有り、カテゴリなし、価格なし
	} else if freewords != "" && (category_first == "99" && category_second == "99" && category_third == "99") && (price_min == 0 && price_max == 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND title like ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			"%"+freewords+"%",
		)
		if err != nil {
			log.Fatalln(err)
		}

		//フリーワードなし、カテゴリ有り、価格有り
	} else if freewords == "" && (category_first != "99" || category_second != "99" || category_third != "99") && (price_min != 0 || price_max != 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?) AND price >= ? AND price <= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			price_min, price_max,
		)
		if err != nil {
			log.Fatalln(err)
		}
		//フリーワードなし、カテゴリ有り、価格なし
	} else if freewords == "" && (category_first != "99" || category_second != "99" || category_third != "99") && (price_min == 0 && price_max == 100000000) {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND category_first IN(?, ?, ?)  OR category_second IN(?, ?, ?) OR category_third IN(?, ?, ?)`

		rows, err = Db.Query(
			cmd,
			u.ID,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
			category_first, category_second, category_third,
		)
		if err != nil {
			log.Fatalln(err)
		}

		//フリーワードなし、カテゴリなし、価格あり
	} else if freewords == "" && (category_first == "99" && category_second == "99" && category_third == "99") && (price_min != 0 || price_max != 100000000) {
		fmt.Println("price")
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ? AND price <= ? AND price >= ?`

		rows, err = Db.Query(
			cmd,
			u.ID,
			price_min, price_max,
		)
	} else {
		cmd = `SELECT id, user_id, created_at, photo_url, title, price FROM items
		WHERE user_id != ?`
		rows, err = Db.Query(
			cmd,
			u.ID,
		)
		if err != nil {
			log.Fatalln(err)
		}
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
	cmd := `SELECT id, user_id, created_at, photo_url, title, price FROM items
	WHERE id in (SELECT item_id FROM items WHERE id =?)`
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
	cmd := `update items set photo_url = ?, title = ?, content = ?, category_first= ? ,category_second = ?, category_third = ?, price = ? WHERE id = ?`

	_, err = Db.Exec(cmd, i.PhotoURL, i.Title, i.Content, i.CategoryFirst, i.CategorySecond, i.CategoryThird, i.Price, i.ID)

	return err

}

func (i *Item) DeleteItem() error {
	cmd := `delete FROM items WHERE id =?`

	_, err = Db.Exec(cmd, i.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
