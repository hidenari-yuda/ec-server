package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hidenari-yuda/todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		fmt.Println(err)
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)

	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}

func profile(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, user, "layout", "private_navbar", "profile")
	}
}

func profileEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		u, err := models.GetUser(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, u, "layout", "private_navbar", "profile_edit")
	}
}

func profileUpdate(w http.ResponseWriter, r *http.Request, id int) {
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
		name, email, phone, department, position := r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("phone"), r.PostFormValue("department"), r.PostFormValue("position")
		u := &models.User{ID: id, Name: name, Email: email, Phone: phone, Department: department, Position: position}
		if err := u.UpdateUser(); err != nil {
			log.Println(err)
		} else {
			http.Redirect(w, r, "/profile", 302)
		}
	}
}

func aboutus(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "aboutus")
}

func contact(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "contact")
}
