package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hidenari-yuda/ec-server/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/onsale", 302)
	}
}

func onsale(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		items, _ := user.GetItemsByOthers()
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "onsale")
	}
}

func onsaleSelect(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		i, _ := models.GetItem(id)
		//saleUser, _ := models.GetUser(item.UserID)
		//user.NickName = saleUser.NickName
		generateHTML(w, i, "layout", "private_navbar", "onsale_item")
	}
}

func purchase(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		i, _ := models.GetItem(id)
		//saleUser, _ := models.GetUser(item.UserID)
		//user.NickName = saleUser.NickName
		generateHTML(w, i, "layout", "private_navbar", "purchase")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		items, _ := user.GetItemsByUser()
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func itemSort(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		r.ParseForm()
		items, _ := user.GetItemsBySort(r.PostFormValue("sort"))
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func onsaleSort(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		r.ParseForm()
		items, _ := user.GetOnsaleBySort(r.PostFormValue("sort"))
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "onsale")
	}
}

func itemSortByCategory(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		r.ParseForm()
		items, _ := user.GetItemsByCategory(
			r.PostFormValue("category_first"),
			r.PostFormValue("category_second"),
			r.PostFormValue("category_third"),
		)
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func onsaleSortByCategory(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		r.ParseForm()
		items, _ := user.GetOnsaleByCategory(
			r.PostFormValue("category_first"),
			r.PostFormValue("category_second"),
			r.PostFormValue("category_third"),
		)
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "onsale")
	}
}

func itemSortByFreeWords(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		r.ParseForm()

		min := r.PostFormValue("price_min")
		minInt, _ := strconv.Atoi(min)

		max := r.PostFormValue("price_max")
		maxInt, _ := strconv.Atoi(max)

		items, _ := user.GetItemsByFreeWords(
			r.PostFormValue("freewords"),
			r.PostFormValue("category_first"),
			r.PostFormValue("category_second"),
			r.PostFormValue("category_third"),
			minInt,
			maxInt,
		)

		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func onsaleSortByFreeWords(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		r.ParseForm()
		min := r.PostFormValue("price_min")
		minInt, _ := strconv.Atoi(min)
		max := r.PostFormValue("price_max")
		maxInt, _ := strconv.Atoi(max)

		items, _ := user.GetOnsaleByFreeWords(
			r.PostFormValue("freewords"),
			r.PostFormValue("category_first"),
			r.PostFormValue("category_second"),
			r.PostFormValue("category_third"),
			minInt,
			maxInt,
		)
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "onsale")
	}
}

func itemNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "item_new")
	}

}

func itemSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		photo_url, title, content, category_first, category_second, category_third, price := r.PostFormValue("photo_url"), r.PostFormValue("title"), r.PostFormValue("content"), r.PostFormValue("category_first"), r.PostFormValue("category_second"), r.PostFormValue("category_third"), r.PostFormValue("price")
		priceInt, _ := strconv.Atoi(price)
		if err := user.CreateItem(photo_url, title, content, category_first, category_second, category_third, priceInt); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/items", 302)
	}
}

func itemEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		i, err := models.GetItem(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, i, "layout", "private_navbar", "item_edit")
	}
}

func itemUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err = sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		photo_url, title, content, category_first, category_second, category_third, price := r.PostFormValue("photo_url"), r.PostFormValue("title"), r.PostFormValue("content"), r.PostFormValue("category_first"), r.PostFormValue("category_second"), r.PostFormValue("category_third"), r.PostFormValue("price")
		priceInt, _ := strconv.Atoi(price)
		i := &models.Item{ID: id, PhotoURL: photo_url, Title: title, Content: content, CategoryFirst: category_first, CategorySecond: category_second, CategoryThird: category_third, Price: priceInt}
		if err := i.UpdateItem(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/items", 302)
	}
}

func itemDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err = sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		i, err := models.GetItem(id)
		if err != nil {
			log.Println(err)
		}
		if err := i.DeleteItem(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/items", 302)
	}

}
