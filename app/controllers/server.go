package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	"github.com/hidenari-yuda/ec-server/app/models"
	"github.com/hidenari-yuda/ec-server/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/items/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

var validPathProfile = regexp.MustCompile("^/profile/(edit|update)/([0-9]+)$")

func parseURLProfile(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPathProfile.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

var validPathChat = regexp.MustCompile("^/chat/([0-9]+)$")

func parseURLChat(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPathChat.FindStringSubmatch(r.URL.Path)
		if q == nil {

			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[1])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

var validPathChatCRUD = regexp.MustCompile("^/chat/(save | edit |update)/([0-9]+)$")

func parseURLChatCRUD(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPathChatCRUD.FindStringSubmatch(r.URL.Path)
		if q == nil {

			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fmt.Println(qi)
		fn(w, r, qi)
	}
}

var validPathChatGroup = regexp.MustCompile("^/group/(edit|update|delete)/([0-9]+)$")

func parseURLChatGroup(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPathChatGroup.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

var validPathOnsale = regexp.MustCompile("^/onsale/([0-9]+)$")

func parseURLOnsale(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPathOnsale.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[1])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/profile/edit/", parseURLProfile(profileEdit))
	http.HandleFunc("/profile/update/", parseURLProfile(profileUpdate))
	http.HandleFunc("/items", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/items/new", itemNew)
	http.HandleFunc("/items/save", itemSave)
	http.HandleFunc("/items/edit/", parseURL(itemEdit))
	http.HandleFunc("/items/update/", parseURL(itemUpdate))
	http.HandleFunc("/items/delete/", parseURL(itemDelete))
	http.HandleFunc("/sort", itemSort)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/aboutus", aboutus)
	http.HandleFunc("/chat/", parseURLChat(chat))
	http.HandleFunc("/chat/save", chatSave)
	http.HandleFunc("/chat/delete/", parseURLChatCRUD(chatDelete))
	http.HandleFunc("/group", chatGroup)
	http.HandleFunc("/group/save", chatGroupSave)
	http.HandleFunc("/group/delete/", parseURLChatGroup(chatGroupDelete))
	http.HandleFunc("/onsale", onsale)
	http.HandleFunc("/onsale/", parseURLOnsale(onsaleSelect))

	//port := os.Getenv("PORT")
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
