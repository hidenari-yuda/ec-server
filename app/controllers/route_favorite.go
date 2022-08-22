package controllers

import (
	"log"
	"net/http"
	"strconv"
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
		//favorites_items, _ := models.GetItemsByFavorites(1, 2)
		//user.Items = favorites_items
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
		id := r.PostFormValue("item_id")
		idInt, err := strconv.Atoi(id)
		if err := user.CreateFavorites(idInt); err != nil {
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
		i, err := models.GetFavorites(id)
		if err != nil {
			log.Println(err)
		}
		if err := i.DeleteFavorites(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/items", 302)
	}

}
