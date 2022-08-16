package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hidenari-yuda/ec-server/app/models"
)

func favorites(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		items, _ := models.GetItemsByFavorites(1)
		user.Items = items
		generateHTML(w, user, "layout", "private_navbar", "favorites")
	}
}

func favoritesSave(w http.ResponseWriter, r *http.Request) {
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

func favoritesDelete(w http.ResponseWriter, r *http.Request, id int) {
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
