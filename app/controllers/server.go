package controllers

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/items/sort", itemSort)
	http.HandleFunc("/onsale/sort", onsaleSort)
	http.HandleFunc("/items/category", itemSortByCategory)
	http.HandleFunc("/onsale/category", onsaleSortByCategory)
	http.HandleFunc("/items/freewords", itemSortByFreeWords)
	http.HandleFunc("/onsale/freewords", onsaleSortByFreeWords)
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
	http.HandleFunc("/purchase/", parseURLPurchase(purchase))
	http.HandleFunc("/favorites", favorites)
	http.HandleFunc("/favorites/save/", favoritesSave)
	http.HandleFunc("/favorites/delete/", parseURLFavorites(favoritesDelete))

	//port := os.Getenv("PORT")
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
